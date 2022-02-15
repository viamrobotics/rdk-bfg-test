// Package status define the status structures of a robot.
package status

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.viam.com/rdk/component/arm"
	"go.viam.com/rdk/component/base"
	"go.viam.com/rdk/component/board"
	"go.viam.com/rdk/component/camera"
	"go.viam.com/rdk/component/gantry"
	"go.viam.com/rdk/component/gripper"
	"go.viam.com/rdk/component/input"
	"go.viam.com/rdk/component/motor"
	"go.viam.com/rdk/component/sensor"
	"go.viam.com/rdk/component/servo"
	commonpb "go.viam.com/rdk/proto/api/common/v1"
	pb "go.viam.com/rdk/proto/api/v1"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot"
)

// Create constructs a new up to date status from the given robot.
// The operation can take time and be expensive, so it can be cancelled by the
// given context.
func Create(ctx context.Context, r robot.Robot) (*pb.Status, error) {
	var err error
	var status pb.Status

	status.Services = make(map[string]bool)

	// manually refresh all remotes to get an up to date status
	for _, name := range r.RemoteNames() {
		remote, ok := r.RemoteByName(name)
		if !ok {
			continue
		}
		if refresher, ok := remote.(robot.Refresher); ok {
			if err := refresher.Refresh(ctx); err != nil {
				return nil, errors.Wrapf(err, "error refreshing remote %q", name)
			}
		}
	}

	for _, name := range r.ResourceNames() {
		switch {
		case name.Subtype == arm.Subtype:
			if status.Arms == nil {
				status.Arms = make(map[string]*pb.ArmStatus)
			}
			raw, ok := r.ResourceByName(name)
			if !ok {
				return nil, errors.New("should be impossible")
			}

			arm, ok := raw.(arm.Arm)
			if !ok {
				return nil, errors.New("should be impossible")
			}

			armStatus := &pb.ArmStatus{}

			gridPosition, err := arm.GetEndPosition(ctx)
			if err != nil {
				return nil, err
			}
			if gridPosition != nil {
				armStatus.GridPosition = &pb.Pose{
					X:     gridPosition.X,
					Y:     gridPosition.Y,
					Z:     gridPosition.Z,
					OX:    gridPosition.OX,
					OY:    gridPosition.OY,
					OZ:    gridPosition.OZ,
					Theta: gridPosition.Theta,
				}
			}

			jointPositions, err := arm.GetJointPositions(ctx)
			if err != nil {
				return nil, err
			}
			if jointPositions != nil {
				armStatus.JointPositions = &pb.JointPositions{
					Degrees: jointPositions.Degrees,
				}
			}

			status.Arms[name.Name] = armStatus
		case name.Subtype == gantry.Subtype:
			if status.Gantries == nil {
				status.Gantries = make(map[string]*pb.GantryStatus)
			}
			raw, ok := r.ResourceByName(name)
			if !ok {
				return nil, errors.New("should be impossible")
			}

			g, ok := raw.(gantry.Gantry)
			if !ok {
				return nil, errors.New("should be impossible")
			}

			gantryStatus := &pb.GantryStatus{}

			gantryStatus.Positions, err = g.GetPosition(ctx)
			if err != nil {
				return nil, err
			}

			gantryStatus.Lengths, err = g.GetLengths(ctx)
			if err != nil {
				return nil, err
			}

			status.Gantries[name.Name] = gantryStatus
		case name.ResourceType == resource.ResourceTypeService:
			status.Services[name.Subtype.String()] = true
		}
	}

	if names := gripper.NamesFromRobot(r); len(names) != 0 {
		status.Grippers = make(map[string]bool, len(names))
		for _, name := range names {
			status.Grippers[name] = true
		}
	}

	if names := base.NamesFromRobot(r); len(names) != 0 {
		status.Bases = make(map[string]bool, len(names))
		for _, name := range names {
			status.Bases[name] = true
		}
	}

	if names := board.NamesFromRobot(r); len(names) != 0 {
		status.Boards = make(map[string]*commonpb.BoardStatus, len(names))
		for _, name := range names {
			board, err := board.FromRobot(r, name)
			if err != nil {
				continue
			}
			boardStatus, err := board.Status(ctx)
			if err != nil {
				return nil, err
			}
			status.Boards[name] = boardStatus
		}
	}

	if names := camera.NamesFromRobot(r); len(names) != 0 {
		status.Cameras = make(map[string]bool, len(names))
		for _, name := range names {
			status.Cameras[name] = true
		}
	}

	if names := sensor.NamesFromRobot(r); len(names) != 0 {
		status.Sensors = make(map[string]*pb.SensorStatus, len(names))
		for _, name := range names {
			status.Sensors[name] = &pb.SensorStatus{
				Type: "sensor",
			}
		}
	}

	if names := servo.NamesFromRobot(r); len(names) != 0 {
		status.Servos = make(map[string]*pb.ServoStatus, len(names))
		for _, name := range names {
			x, err := servo.FromRobot(r, name)
			if err != nil {
				return nil, err
			}
			current, err := x.GetPosition(ctx)
			if err != nil {
				return nil, err
			}
			status.Servos[name] = &pb.ServoStatus{
				Angle: uint32(current),
			}
		}
	}

	if names := r.MotorNames(); len(names) != 0 {
		status.Motors = make(map[string]*pb.MotorStatus, len(names))
		for _, name := range names {
			x, ok := r.MotorByName(name)
			if !ok {
				return nil, fmt.Errorf("motor %q not found", name)
			}
			isOn, err := x.IsPowered(ctx)
			if err != nil {
				return nil, err
			}
			position, err := x.GetPosition(ctx)
			if err != nil {
				return nil, err
			}
			features, err := x.GetFeatures(ctx)
			if err != nil {
				return nil, err
			}
			positionSupported := features[motor.PositionReporting]
			var str *structpb.Struct
			status.Motors[name] = &pb.MotorStatus{
				On:                isOn,
				Position:          position,
				PositionSupported: positionSupported,
				PidConfig:         str,
			}
		}
	}

	if names := input.NamesFromRobot(r); len(names) != 0 {
		status.InputControllers = make(map[string]*pb.InputControllerStatus, len(names))
		for _, name := range names {
			controller, err := input.FromRobot(r, name)
			if err != nil {
				return nil, err
			}
			eventsIn, err := controller.GetEvents(ctx)
			if err != nil {
				return nil, err
			}
			resp := &pb.InputControllerStatus{}
			for _, eventIn := range eventsIn {
				resp.Events = append(resp.Events, &pb.InputControllerEvent{
					Time:    timestamppb.New(eventIn.Time),
					Event:   string(eventIn.Event),
					Control: string(eventIn.Control),
					Value:   eventIn.Value,
				})
			}
			status.InputControllers[name] = resp
		}
	}

	if names := r.FunctionNames(); len(names) != 0 {
		status.Functions = make(map[string]bool, len(names))
		for _, name := range names {
			status.Functions[name] = true
		}
	}

	return &status, nil
}