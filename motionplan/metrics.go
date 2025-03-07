package motionplan

import (
	"math"

	"go.viam.com/rdk/referenceframe"
	spatial "go.viam.com/rdk/spatialmath"
	"go.viam.com/rdk/utils"
)

const orientationDistanceScaling = 10.

// StateMetric are functions which, given a State, produces some score. Lower is better.
// This is used for gradient descent to converge upon a goal pose, for example.
type StateMetric func(*State) float64

// SegmentMetric are functions which produce some score given an Segment. Lower is better.
// This is used to sort produced IK solutions by goodness, for example.
type SegmentMetric func(*Segment) float64

// NewZeroMetric always returns zero as the distance between two points.
func NewZeroMetric() StateMetric {
	return func(from *State) float64 { return 0 }
}

type combinableStateMetric struct {
	metrics []StateMetric
}

func (m *combinableStateMetric) combinedDist(input *State) float64 {
	dist := 0.
	for _, metric := range m.metrics {
		dist += metric(input)
	}
	return dist
}

// CombineMetrics will take a variable number of Metrics and return a new Metric which will combine all given metrics into one, summing
// their distances.
func CombineMetrics(metrics ...StateMetric) StateMetric {
	cm := &combinableStateMetric{metrics: metrics}
	return cm.combinedDist
}

// orientDist returns the arclength between two orientations in degrees.
func orientDist(o1, o2 spatial.Orientation) float64 {
	return utils.RadToDeg(spatial.QuatToR4AA(spatial.OrientationBetween(o1, o2).Quaternion()).Theta)
}

// orientDistToRegion will return a function which will tell you how far the unit sphere component of an orientation
// vector is from a region defined by a point and an arclength around it. The theta value of OV is disregarded.
// This is useful, for example, in defining the set of acceptable angles of attack for writing on a whiteboard.
func orientDistToRegion(goal spatial.Orientation, alpha float64) func(spatial.Orientation) float64 {
	ov1 := goal.OrientationVectorRadians()
	return func(o spatial.Orientation) float64 {
		ov2 := o.OrientationVectorRadians()
		acosInput := ov1.OX*ov2.OX + ov1.OY*ov2.OY + ov1.OZ*ov2.OZ

		// Account for floating point issues
		if acosInput > 1.0 {
			acosInput = 1.0
		}
		if acosInput < -1.0 {
			acosInput = -1.0
		}
		dist := math.Acos(acosInput)
		return math.Max(0, dist-alpha)
	}
}

// NewSquaredNormMetric is the default distance function between two poses to be used for gradient descent.
func NewSquaredNormMetric(goal spatial.Pose) StateMetric {
	weightedSqNormDist := func(query *State) float64 {
		delta := spatial.PoseDelta(goal, query.Position)
		// Increase weight for orientation since it's a small number
		return delta.Point().Norm2() + spatial.QuatToR3AA(delta.Orientation().Quaternion()).Mul(orientationDistanceScaling).Norm2()
	}
	return weightedSqNormDist
}

// NewPoseFlexOVMetric will provide a distance function which will converge on a pose with an OV within an arclength of `alpha`
// of the ov of the goal given.
func NewPoseFlexOVMetric(goal spatial.Pose, alpha float64) StateMetric {
	oDistFunc := orientDistToRegion(goal.Orientation(), alpha)
	return func(state *State) float64 {
		pDist := state.Position.Point().Distance(goal.Point())
		oDist := oDistFunc(state.Position.Orientation())
		return pDist*pDist + oDist*oDist
	}
}

// NewPositionOnlyMetric returns a Metric that reports the point-wise distance between two poses without regard for orientation.
// This is useful for scenarios where there are not enough DOF to control orientation, but arbitrary spatial points may
// still be arrived at.
func NewPositionOnlyMetric(goal spatial.Pose) StateMetric {
	return func(state *State) float64 {
		pDist := state.Position.Point().Distance(goal.Point())
		return pDist * pDist
	}
}

// JointMetric is a metric which will sum the squared differences in each input from start to end.
func JointMetric(segment *Segment) float64 {
	jScore := 0.
	for i, f := range segment.StartConfiguration {
		jScore += math.Abs(f.Value - segment.EndConfiguration[i].Value)
	}
	return jScore
}

// L2InputMetric is a metric which will return a L2 norm of the StartConfiguration and EndConfiguration in an arc input.
func L2InputMetric(segment *Segment) float64 {
	return referenceframe.InputsL2Distance(segment.StartConfiguration, segment.EndConfiguration)
}

// TODO(RSDK-2557): Writing a PenetrationDepthMetric will allow cbirrt to path along the sides of obstacles rather than terminating
// the RRT tree when an obstacle is hit
