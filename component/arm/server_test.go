package arm_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"go.viam.com/test"

	"go.viam.com/rdk/component/arm"
	commonpb "go.viam.com/rdk/proto/api/common/v1"
	pb "go.viam.com/rdk/proto/api/component/v1"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/subtype"
	"go.viam.com/rdk/testutils/inject"
)

func newServer() (pb.ArmServiceServer, *inject.Arm, *inject.Arm, error) {
	injectArm := &inject.Arm{}
	injectArm2 := &inject.Arm{}
	arms := map[resource.Name]interface{}{
		arm.Named(testArmName): injectArm,
		arm.Named(failArmName): injectArm2,
		arm.Named(fakeArmName): "notArm",
	}
	armSvc, err := subtype.New(arms)
	if err != nil {
		return nil, nil, nil, err
	}
	return arm.NewServer(armSvc), injectArm, injectArm2, nil
}

func TestServer(t *testing.T) {
	armServer, injectArm, injectArm2, err := newServer()
	test.That(t, err, test.ShouldBeNil)

	var (
		capArmPos      *commonpb.Pose
		capArmJointPos *pb.ArmJointPositions
	)

	pose1 := &commonpb.Pose{X: 1, Y: 2, Z: 3}
	positionDegs1 := &pb.ArmJointPositions{Degrees: []float64{1.0, 2.0, 3.0}}
	injectArm.GetEndPositionFunc = func(ctx context.Context) (*commonpb.Pose, error) {
		return pose1, nil
	}
	injectArm.GetJointPositionsFunc = func(ctx context.Context) (*pb.ArmJointPositions, error) {
		return positionDegs1, nil
	}
	injectArm.MoveToPositionFunc = func(ctx context.Context, ap *commonpb.Pose) error {
		capArmPos = ap
		return nil
	}

	injectArm.MoveToJointPositionsFunc = func(ctx context.Context, jp *pb.ArmJointPositions) error {
		capArmJointPos = jp
		return nil
	}

	pose2 := &commonpb.Pose{X: 4, Y: 5, Z: 6}
	positionDegs2 := &pb.ArmJointPositions{Degrees: []float64{4.0, 5.0, 6.0}}
	injectArm2.GetEndPositionFunc = func(ctx context.Context) (*commonpb.Pose, error) {
		return nil, errors.New("can't get pose")
	}
	injectArm2.GetJointPositionsFunc = func(ctx context.Context) (*pb.ArmJointPositions, error) {
		return nil, errors.New("can't get joint positions")
	}
	injectArm2.MoveToPositionFunc = func(ctx context.Context, ap *commonpb.Pose) error {
		capArmPos = ap
		return errors.New("can't move to pose")
	}

	injectArm2.MoveToJointPositionsFunc = func(ctx context.Context, jp *pb.ArmJointPositions) error {
		capArmJointPos = jp
		return errors.New("can't move to joint positions")
	}

	t.Run("arm position", func(t *testing.T) {
		_, err := armServer.GetEndPosition(context.Background(), &pb.ArmServiceGetEndPositionRequest{Name: missingArmName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no arm")

		_, err = armServer.GetEndPosition(context.Background(), &pb.ArmServiceGetEndPositionRequest{Name: fakeArmName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "not an arm")

		resp, err := armServer.GetEndPosition(context.Background(), &pb.ArmServiceGetEndPositionRequest{Name: testArmName})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, resp.Pose.String(), test.ShouldResemble, pose1.String())

		_, err = armServer.GetEndPosition(context.Background(), &pb.ArmServiceGetEndPositionRequest{Name: failArmName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't get pose")
	})

	//nolint:dupl
	t.Run("move to position", func(t *testing.T) {
		_, err = armServer.MoveToPosition(context.Background(), &pb.ArmServiceMoveToPositionRequest{Name: missingArmName, Pose: pose2})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no arm")

		_, err := armServer.MoveToPosition(context.Background(), &pb.ArmServiceMoveToPositionRequest{Name: testArmName, Pose: pose2})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, capArmPos.String(), test.ShouldResemble, pose2.String())

		_, err = armServer.MoveToPosition(context.Background(), &pb.ArmServiceMoveToPositionRequest{Name: failArmName, Pose: pose1})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't move to pose")
		test.That(t, capArmPos.String(), test.ShouldResemble, pose1.String())
	})

	t.Run("arm joint position", func(t *testing.T) {
		_, err := armServer.GetJointPositions(context.Background(), &pb.ArmServiceGetJointPositionsRequest{Name: missingArmName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no arm")

		resp, err := armServer.GetJointPositions(context.Background(), &pb.ArmServiceGetJointPositionsRequest{Name: testArmName})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, resp.PositionDegs.String(), test.ShouldResemble, positionDegs1.String())

		_, err = armServer.GetJointPositions(context.Background(), &pb.ArmServiceGetJointPositionsRequest{Name: failArmName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't get joint positions")
	})

	//nolint:dupl
	t.Run("move to joint position", func(t *testing.T) {
		_, err = armServer.MoveToJointPositions(
			context.Background(),
			&pb.ArmServiceMoveToJointPositionsRequest{Name: missingArmName, PositionDegs: positionDegs2},
		)
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no arm")

		_, err := armServer.MoveToJointPositions(
			context.Background(),
			&pb.ArmServiceMoveToJointPositionsRequest{Name: testArmName, PositionDegs: positionDegs2},
		)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, capArmJointPos.String(), test.ShouldResemble, positionDegs2.String())

		_, err = armServer.MoveToJointPositions(
			context.Background(),
			&pb.ArmServiceMoveToJointPositionsRequest{Name: failArmName, PositionDegs: positionDegs1},
		)
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't move to joint positions")
		test.That(t, capArmJointPos.String(), test.ShouldResemble, positionDegs1.String())
	})
}