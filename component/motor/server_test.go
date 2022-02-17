package motor_test

import (
	"context"
	"errors"
	"testing"

	"go.viam.com/test"

	"go.viam.com/rdk/component/motor"
	pb "go.viam.com/rdk/proto/api/component/v1"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/subtype"
	"go.viam.com/rdk/testutils/inject"
)

func newServer() (pb.MotorServiceServer, *inject.Motor, *inject.Motor, error) {
	injectMotor1 := &inject.Motor{}
	injectMotor2 := &inject.Motor{}

	resourceMap := map[resource.Name]interface{}{
		motor.Named("workingMotor"): injectMotor1,
		motor.Named("failingMotor"): injectMotor2,
		motor.Named("notAMotor"):    "not a motor",
	}

	injectSvc, err := subtype.New(resourceMap)
	if err != nil {
		return nil, nil, nil, err
	}
	return motor.NewServer(injectSvc), injectMotor1, injectMotor2, nil
}

//nolint:dupl
func TestSetPower(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceSetPowerRequest{Name: "notAMotor"}
	resp, err := motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.SetPowerFunc = func(ctx context.Context, powerPct float64) error {
		return errors.New("set power failed")
	}
	req = pb.MotorServiceSetPowerRequest{Name: "failingMotor", PowerPct: 0.5}
	resp, err = motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.SetPowerFunc = func(ctx context.Context, powerPct float64) error {
		return nil
	}
	req = pb.MotorServiceSetPowerRequest{Name: "workingMotor", PowerPct: 0.5}
	resp, err = motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

//nolint:dupl
func TestGoFor(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGoForRequest{Name: "notAMotor"}
	resp, err := motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GoForFunc = func(ctx context.Context, rpm, rotations float64) error {
		return errors.New("go for failed")
	}
	req = pb.MotorServiceGoForRequest{Name: "failingMotor", Rpm: 42.0, Revolutions: 42.1}
	resp, err = motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GoForFunc = func(ctx context.Context, rpm, rotations float64) error {
		return nil
	}
	req = pb.MotorServiceGoForRequest{Name: "workingMotor", Rpm: 42.0, Revolutions: 42.1}
	resp, err = motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestPosition(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGetPositionRequest{Name: "notAMotor"}
	resp, err := motorServer.GetPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GetPositionFunc = func(ctx context.Context) (float64, error) {
		return 0, errors.New("position unavailable")
	}
	req = pb.MotorServiceGetPositionRequest{Name: "failingMotor"}
	resp, err = motorServer.GetPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GetPositionFunc = func(ctx context.Context) (float64, error) {
		return 42.0, nil
	}
	req = pb.MotorServiceGetPositionRequest{Name: "workingMotor"}
	resp, err = motorServer.GetPosition(context.Background(), &req)
	test.That(t, resp.GetPosition(), test.ShouldEqual, 42.0)
	test.That(t, err, test.ShouldBeNil)
}

func TestGetFeatures(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGetFeaturesRequest{Name: "notAMotor"}
	resp, err := motorServer.GetFeatures(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GetFeaturesFunc = func(ctx context.Context) (map[motor.Feature]bool, error) {
		return nil, errors.New("unable to get supported features")
	}
	req = pb.MotorServiceGetFeaturesRequest{Name: "failingMotor"}
	resp, err = motorServer.GetFeatures(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GetFeaturesFunc = func(ctx context.Context) (map[motor.Feature]bool, error) {
		return map[motor.Feature]bool{
			motor.PositionReporting: true,
		}, nil
	}
	req = pb.MotorServiceGetFeaturesRequest{Name: "workingMotor"}
	resp, err = motorServer.GetFeatures(context.Background(), &req)
	test.That(t, resp.GetPositionReporting(), test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
}

func TestStop(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceStopRequest{Name: "notAMotor"}
	resp, err := motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.StopFunc = func(ctx context.Context) error {
		return errors.New("stop failed")
	}
	req = pb.MotorServiceStopRequest{Name: "failingMotor"}
	resp, err = motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.StopFunc = func(ctx context.Context) error {
		return nil
	}
	req = pb.MotorServiceStopRequest{Name: "workingMotor"}
	resp, err = motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestIsOn(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceIsPoweredRequest{Name: "notAMotor"}
	resp, err := motorServer.IsPowered(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.IsPoweredFunc = func(ctx context.Context) (bool, error) {
		return false, errors.New("could not determine if motor is on")
	}
	req = pb.MotorServiceIsPoweredRequest{Name: "failingMotor"}
	resp, err = motorServer.IsPowered(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.IsPoweredFunc = func(ctx context.Context) (bool, error) {
		return true, nil
	}
	req = pb.MotorServiceIsPoweredRequest{Name: "workingMotor"}
	resp, err = motorServer.IsPowered(context.Background(), &req)
	test.That(t, resp.GetIsOn(), test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
}

//nolint:dupl
func TestGoTo(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGoToRequest{Name: "notAMotor"}
	resp, err := motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GoToFunc = func(ctx context.Context, rpm, position float64) error {
		return errors.New("go to failed")
	}
	req = pb.MotorServiceGoToRequest{Name: "failingMotor", Rpm: 20.0, PositionRevolutions: 2.5}
	resp, err = motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GoToFunc = func(ctx context.Context, rpm, position float64) error {
		return nil
	}
	req = pb.MotorServiceGoToRequest{Name: "workingMotor", Rpm: 20.0, PositionRevolutions: 2.5}
	resp, err = motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

//nolint:dupl
func TestResetZeroPosition(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceResetZeroPositionRequest{Name: "notAMotor"}
	resp, err := motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.ResetZeroPositionFunc = func(ctx context.Context, offset float64) error {
		return errors.New("set to zero failed")
	}
	req = pb.MotorServiceResetZeroPositionRequest{Name: "failingMotor", Offset: 1.1}
	resp, err = motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.ResetZeroPositionFunc = func(ctx context.Context, offset float64) error {
		return nil
	}
	req = pb.MotorServiceResetZeroPositionRequest{Name: "workingMotor", Offset: 1.1}
	resp, err = motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}