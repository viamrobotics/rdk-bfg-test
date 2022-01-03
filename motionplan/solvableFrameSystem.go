package motionplan

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"github.com/edaniels/golog"
	"go.uber.org/multierr"

	frame "go.viam.com/rdk/referenceframe"
	spatial "go.viam.com/rdk/spatialmath"
)

// SolvableFrameSystem wraps a FrameSystem to allow solving between frames of the frame system.
// Note that this needs to live in motionplan, not referenceframe, to avoid circular dependencies.
type SolvableFrameSystem struct {
	frame.FrameSystem
	logger golog.Logger
	mpFunc func(frame.Frame, int, golog.Logger) (MotionPlanner, error)
}

// NewSolvableFrameSystem will create a new solver for a frame system.
func NewSolvableFrameSystem(fs frame.FrameSystem, logger golog.Logger) *SolvableFrameSystem {
	return &SolvableFrameSystem{FrameSystem: fs, logger: logger}
}

// SolvePose will take a set of starting positions, a goal frame, a frame to solve for, and a pose. The function will
// then try to path plan the full frame system such that the solveFrame has the goal pose from the perspective of the goalFrame.
// For example, if a world system has a gripper attached to an arm attached to a gantry, and the system was being solved
// to place the gripper at a particular pose in the world, the solveFrame would be the gripper and the goalFrame would be
// the world frame. It will use the default planner options.
func (fss *SolvableFrameSystem) SolvePose(ctx context.Context,
	seedMap map[string][]frame.Input,
	goal spatial.Pose,
	solveFrame, goalFrame frame.Frame,
) ([]map[string][]frame.Input, error) {
	return fss.SolvePoseWithOptions(ctx, seedMap, goal, solveFrame, goalFrame, nil)
}

// SolvePoseWithOptions will take a set of starting positions, a goal frame, a frame to solve for, a pose, and a configurable
// set of PlannerOptions. It will solve the solveFrame to the goal pose with respect to the goal frame using the provided
// planning options.
func (fss *SolvableFrameSystem) SolvePoseWithOptions(ctx context.Context,
	seedMap map[string][]frame.Input,
	goal spatial.Pose,
	solveFrame, goalFrame frame.Frame,
	opt *PlannerOptions,
) ([]map[string][]frame.Input, error) {
	// Default for opt if nil
	if opt == nil {
		opt = NewDefaultPlannerOptions()
	}

	// Get parentage of both frames. This will also verify the frames are in the frame system
	sFrames, err := fss.TracebackFrame(solveFrame)
	if err != nil {
		return nil, err
	}
	gFrames, err := fss.TracebackFrame(goalFrame)
	if err != nil {
		return nil, err
	}
	frames := uniqInPlaceSlice(append(sFrames, gFrames...))

	// Create a frame to solve for, and an IK solver with that frame.
	sf := &solverFrame{solveFrame.Name() + "_" + goalFrame.Name(), fss, frames, solveFrame, goalFrame}
	if len(sf.DoF()) == 0 {
		return nil, errors.New("solver frame has no degrees of freedom, cannot perform inverse kinematics")
	}
	var planner MotionPlanner
	if fss.mpFunc != nil {
		planner, err = fss.mpFunc(sf, runtime.NumCPU()/2, fss.logger)
	} else {
		planner, err = NewCBiRRTMotionPlanner(sf, runtime.NumCPU()/2, fss.logger)
	}
	if err != nil {
		return nil, err
	}

	seed := sf.mapToSlice(seedMap)

	// Solve for the goal position
	resultSlices, err := planner.Plan(ctx, spatial.PoseToProtobuf(goal), seed, opt)
	if err != nil {
		return nil, err
	}
	steps := make([]map[string][]frame.Input, 0, len(resultSlices))
	for _, resultSlice := range resultSlices {
		steps = append(steps, sf.sliceToMap(resultSlice))
	}

	return steps, nil
}

// SetPlannerGen sets the function which is used to create the motion planner to solve a requested plan.
// A SolvableFrameSystem wraps a complete frame system, and will make solverFrames on the fly to solve for. These
// solverFrames are used to create the planner here.
func (fss *SolvableFrameSystem) SetPlannerGen(mpFunc func(frame.Frame, int, golog.Logger) (MotionPlanner, error)) {
	fss.mpFunc = mpFunc
}

// solverFrames are meant to be ephemerally created each time a frame system solution is created, and fulfills the
// Frame interface so that it can be passed to inverse kinematics.
type solverFrame struct {
	name       string
	fss        *SolvableFrameSystem
	frames     []frame.Frame
	solveFrame frame.Frame
	goalFrame  frame.Frame
}

// Name returns the name of the solver referenceframe.
func (sf *solverFrame) Name() string {
	return sf.name
}

// Transform returns the pose between the two frames of this solver for a given set of inputs.
func (sf *solverFrame) Transform(inputs []frame.Input) (spatial.Pose, error) {
	if len(inputs) != len(sf.DoF()) {
		return nil, fmt.Errorf("incorrect number of inputs to Transform got %d want %d", len(inputs), len(sf.DoF()))
	}
	return sf.fss.TransformFrame(sf.sliceToMap(inputs), sf.solveFrame, sf.goalFrame)
}

// Volume takes a solverFrame and a list of joint angles in radians and computes the 3D space occupied by each of the
// intermediate frames (if any exist) up to and including the end effector, and eturns a map of frame names to volumes.
// The key for each frame in the map will be the string: "<model_name>:<frame_name>".
func (sf *solverFrame) Volumes(inputs []frame.Input) (map[string]spatial.Volume, error) {
	if len(inputs) != len(sf.DoF()) {
		return nil, errors.New("incorrect number of inputs to transform")
	}
	var errAll error
	inputMap := sf.sliceToMap(inputs)
	volumes := make(map[string]spatial.Volume)
	for _, frame := range sf.frames {
		vols, err := sf.fss.VolumesOfFrame(inputMap, frame, sf.goalFrame)
		if vols == nil {
			// only propagate errors that result in nil volume
			multierr.AppendInto(&errAll, err)
			continue
		}
		for name, vol := range vols {
			volumes[name] = vol
		}
	}
	return volumes, errAll
}

// DoF returns the summed DoF of all frames between the two solver frames.
func (sf *solverFrame) DoF() []frame.Limit {
	var limits []frame.Limit
	for _, frame := range sf.frames {
		limits = append(limits, frame.DoF()...)
	}
	return limits
}

// mapToSlice will flatten a map of inputs into a slice suitable for input to inverse kinematics, by concatenating
// the inputs together in the order of the frames in sf.frames.
func (sf *solverFrame) mapToSlice(inputMap map[string][]frame.Input) []frame.Input {
	var inputs []frame.Input
	for _, frame := range sf.frames {
		inputs = append(inputs, inputMap[frame.Name()]...)
	}
	return inputs
}

func (sf *solverFrame) sliceToMap(inputSlice []frame.Input) map[string][]frame.Input {
	inputs := frame.StartPositions(sf.fss)
	i := 0
	for _, frame := range sf.frames {
		fLen := i + len(frame.DoF())
		inputs[frame.Name()] = inputSlice[i:fLen]
		i = fLen
	}
	return inputs
}

func (sf *solverFrame) MarshalJSON() ([]byte, error) {
	return nil, errors.New("cannot serialize solverFrame")
}

func (sf *solverFrame) AlmostEquals(otherFrame frame.Frame) bool {
	return false
}

// uniqInPlaceSlice will deduplicate the values in a slice using in-place replacement on the slice. This is faster than
// a solution using append().
// This function does not remove anything from the input slice, but it does rearrange the elements.
func uniqInPlaceSlice(s []frame.Frame) []frame.Frame {
	seen := make(map[frame.Frame]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
