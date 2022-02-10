package oneaxis

import (
	"context"
	"errors"
	"testing"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	"go.viam.com/test"

	"go.viam.com/rdk/component/motor"
	"go.viam.com/rdk/component/motor/fake"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	spatial "go.viam.com/rdk/spatialmath"
	"go.viam.com/rdk/testutils/inject"
)

func createFakeMotor() *inject.Motor {
	fakeMotor := &inject.Motor{}

	fakeMotor.PositionSupportedFunc = func(ctx context.Context) (bool, error) {
		return true, nil
	}

	fakeMotor.PositionFunc = func(ctx context.Context) (float64, error) {
		return 1, nil
	}

	fakeMotor.GoForFunc = func(ctx context.Context, rpm float64, revolutions float64) error {
		return nil
	}

	fakeMotor.StopFunc = func(ctx context.Context) error {
		return nil
	}

	fakeMotor.GoFunc = func(ctx context.Context, powerPct float64) error {
		return nil
	}

	return fakeMotor
}

func createFakeBoard() *inject.Board {
	fakeboard := &inject.Board{}

	fakeboard.GetGPIOFunc = func(ctx context.Context, pin string) (bool, error) {
		return true, nil
	}
	return fakeboard
}

func createFakeRobot() *inject.Robot {
	fakerobot := &inject.Robot{}

	fakerobot.MotorByNameFunc = func(name string) (motor.Motor, bool) {
		return &fake.Motor{PositionSupportedFunc: true, GoForfunc: true}, true
	}

	fakerobot.ResourceByNameFunc = func(name resource.Name) (interface{}, bool) {
		return &inject.Board{GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) {
			return true, nil
		}}, true
	}
	return fakerobot
}

func TestValidate(t *testing.T) {
	fakecfg := &AttrConfig{
		Motor:           "x",
		LimitSwitchPins: []string{"1", "2"},
		LengthMm:        0.0,
		Board:           "board",
	}
	err := fakecfg.Validate("path")
	test.That(t, err.Error(), test.ShouldContainSubstring, "each axis needs a non-zero and positive length")

	fakecfg = &AttrConfig{
		LimitSwitchPins: []string{"1"},
		LengthMm:        1.0,
		Board:           "board",
		ReductionRatio:  0.1,
	}
	err = fakecfg.Validate("path")
	test.That(t, err.Error(), test.ShouldContainSubstring, "cannot find motor for gantry")

	fakecfg = &AttrConfig{
		LimitSwitchPins: []string{"1"},
		LengthMm:        1.0,
		ReductionRatio:  0.1,
	}
	err = fakecfg.Validate("path")
	test.That(t, err.Error(), test.ShouldContainSubstring, "cannot find board for gantry")

	fakecfg = &AttrConfig{
		Motor:           "x",
		LimitSwitchPins: []string{"1"},
		LengthMm:        1.0,
		Board:           "board",
		ReductionRatio:  0.0,
	}
	err = fakecfg.Validate("path")
	test.That(t, err.Error(), test.ShouldContainSubstring, "gantry has one limit switch per axis, needs pulley radius to set position limits")

	fakecfg = &AttrConfig{
		Motor:           "x",
		LimitSwitchPins: []string{},
		LengthMm:        1.0,
		Board:           "board",
		ReductionRatio:  0.1,
	}
	err = fakecfg.Validate("path")
	test.That(t, err.Error(), test.ShouldContainSubstring, "each axis needs at least one limit switch pin")

	fakecfg = &AttrConfig{
		Motor:           "x",
		LimitSwitchPins: []string{"1", "2"},
		LengthMm:        1.0,
		Board:           "board",
		ReductionRatio:  0.1,
		Axis: spatial.TranslationConfig{
			X: 1,
			Y: 0,
			Z: 0,
		},
	}

	err = fakecfg.Validate("path")
	test.That(t, err, test.ShouldBeNil)
}

func TestNewOneAxis(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	fakeRobot := createFakeRobot()
	fakecfg := config.Component{Name: "gantry"}
	_, err := newOneAxis(ctx, fakeRobot, fakecfg, logger)
	test.That(t, err.Error(), test.ShouldContainSubstring, "expected *oneaxis.AttrConfig but got <nil>")

	fakecfg = config.Component{
		Name: "gantry",
		ConvertedAttributes: &AttrConfig{
			Motor:           "x",
			LimitSwitchPins: []string{"1", "2"},
			LengthMm:        1.0,
			Board:           "board",
			LimitPinEnabled: true,
			GantryRPM:       float64(300),
		},
	}
	fakegantry, err := newOneAxis(ctx, fakeRobot, fakecfg, logger)
	fakeoneax, ok := fakegantry.(*oneAxis)
	test.That(t, ok, test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, fakeoneax.limitType, test.ShouldEqual, "twoPin")

	fakecfg = config.Component{
		Name: "gantry",
		ConvertedAttributes: &AttrConfig{
			Motor:           "x",
			LimitSwitchPins: []string{"1"},
			LengthMm:        1.0,
			Board:           "board",
			LimitPinEnabled: true,
			ReductionRatio:  0.1,
			GantryRPM:       float64(300),
		},
	}
	fakegantry, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	fakeoneax, ok = fakegantry.(*oneAxis)
	test.That(t, ok, test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, fakeoneax.limitType, test.ShouldEqual, "onePinOneLength")

	fakecfg = config.Component{
		Name: "gantry",
		ConvertedAttributes: &AttrConfig{
			Motor:           "x",
			LimitSwitchPins: []string{"1"},
			LengthMm:        1.0,
			Board:           "board",
			LimitPinEnabled: true,
			GantryRPM:       float64(300),
		},
	}
	fakegantry, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	_, ok = fakegantry.(*oneAxis)
	test.That(t, ok, test.ShouldBeFalse)
	test.That(t, err.Error(), test.ShouldContainSubstring, "gantry with one limit switch per axis needs a reduction ratio defined")

	fakecfg = config.Component{
		Name: "gantry",
		ConvertedAttributes: &AttrConfig{
			LimitSwitchPins: []string{},
			LengthMm:        1.0,
			Board:           "board",
			Motor:           "x",
		},
	}
	_, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	test.That(t, err.Error(), test.ShouldContainSubstring, "encoder currently not supported")

	fakecfg = config.Component{
		Name: "gantry",
		ConvertedAttributes: &AttrConfig{
			Motor:           "x",
			LimitSwitchPins: []string{"1", "2", "3"},
			LengthMm:        1.0,
			Board:           "board",
		},
	}
	_, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	test.That(t, err.Error(), test.ShouldContainSubstring, "invalid gantry type: need 1, 2 or 0 pins per axis, have 3 pins")

	fakeRobot = &inject.Robot{MotorByNameFunc: func(name string) (motor.Motor, bool) { return nil, false }}
	_, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	test.That(t, err.Error(), test.ShouldContainSubstring, "cannot find motor")

	fakeRobot = &inject.Robot{
		MotorByNameFunc:    func(name string) (motor.Motor, bool) { return nil, true },
		ResourceByNameFunc: func(name resource.Name) (interface{}, bool) { return nil, false },
	}

	_, err = newOneAxis(ctx, fakeRobot, fakecfg, logger)
	test.That(t, err.Error(), test.ShouldContainSubstring, "board")
}

func TestHome(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	fakegantry := &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1"},
		limitType:       "onePinOneLength",
	}
	err := fakegantry.Home(ctx)
	test.That(t, err, test.ShouldBeNil)

	fakegantry = &oneAxis{
		motor:     &fake.Motor{PositionSupportedFunc: false},
		limitType: "onePinOneLength",
	}
	err = fakegantry.Home(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry = &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1", "2"},
		limitType:       "twoPin",
	}
	err = fakegantry.Home(ctx)
	test.That(t, err, test.ShouldBeNil)

	fakegantry = &oneAxis{
		motor:     &fake.Motor{PositionSupportedFunc: false},
		limitType: "twoPin",
	}
	err = fakegantry.Home(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry = &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1", "2"},
		limitType:       "encoder",
	}
	err = fakegantry.Home(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, errors.New("err") }}
	err = fakegantry.Home(ctx)
	test.That(t, err, test.ShouldNotBeNil)
}

func TestHomeTwoLimitSwitch(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	fakegantry := &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1", "2"},
	}

	err := fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, fakegantry.positionLimits, test.ShouldResemble, []float64{1, 1})

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, errors.New("err") },
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, nil },
		GoForFunc:             func(ctx context.Context, rpm float64, rotations float64) error { return errors.New("err") },
		StopFunc:              func(ctx context.Context) error { return nil },
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, nil },
		GoForFunc:             func(ctx context.Context, rpm float64, rotations float64) error { return nil },
		StopFunc:              func(ctx context.Context) error { return errors.New("err") },
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, nil },
		GoForFunc:             func(ctx context.Context, rpm float64, rotations float64) error { return errors.New("err") },
		StopFunc:              func(ctx context.Context) error { return nil },
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.board = &inject.Board{
		GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) {
			return true, errors.New("not supported")
		},
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.board = &inject.Board{
		GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) {
			if pin == "1" {
				return true, nil
			}
			return true, errors.New("not supported")
		},
	}
	err = fakegantry.homeTwoLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)
}

func TestHomeOneLimitSwitch(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	fakegantry := &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1"},
		lengthMm:        float64(1),
		reductionRatio:  float64(.1),
	}

	err := fakegantry.homeOneLimSwitch(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, fakegantry.positionLimits, test.ShouldResemble, []float64{1, 2.5915494309189535})

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, errors.New("not supported") },
	}
	err = fakegantry.homeOneLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{
		PositionSupportedFunc: func(ctx context.Context) (bool, error) { return true, nil },
		GoForFunc:             func(ctx context.Context, rpm float64, rotations float64) error { return errors.New("not supported") },
		StopFunc:              func(ctx context.Context) error { return nil },
	}
	err = fakegantry.homeOneLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.board = &inject.Board{
		GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) {
			return true, errors.New("not supported")
		},
	}
	err = fakegantry.homeOneLimSwitch(ctx)
	test.That(t, err, test.ShouldNotBeNil)
}

func TestHomeEncoder(t *testing.T) {
	fakegantry := &oneAxis{}
	ctx := context.Background()
	err := fakegantry.homeEncoder(ctx)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err.Error(), test.ShouldEqual, "encoder currently not supported")
}

func TestTestLimit(t *testing.T) {
	ctx := context.Background()
	fakegantry := &oneAxis{
		limitSwitchPins: []string{"1", "2"},
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		rpm:             float64(300),
		limitHigh:       true,
	}
	pos, err := fakegantry.testLimit(ctx, true)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, pos, test.ShouldEqual, float64(1))
}

func TestLimitHit(t *testing.T) {
	ctx := context.Background()
	fakegantry := &oneAxis{
		limitSwitchPins: []string{"1", "2", "3"},
		board:           createFakeBoard(),
		limitHigh:       true,
	}

	hit, err := fakegantry.limitHit(ctx, true)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, hit, test.ShouldEqual, true)
}

func TestGetPosition(t *testing.T) {
	logger := golog.NewTestLogger(t)
	ctx := context.Background()
	fakegantry := &oneAxis{
		motor: &inject.Motor{
			PositionSupportedFunc: func(ctx context.Context) (bool, error) { return false, nil },
			PositionFunc:          func(ctx context.Context) (float64, error) { return 1, nil },
		},
		board:           createFakeBoard(),
		positionLimits:  []float64{0, 1},
		limitHigh:       true,
		limitSwitchPins: []string{"1", "2"},
		limitType:       switchLimitTypetwoPin,
		logger:          logger,
	}
	positions, err := fakegantry.GetPosition(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, positions, test.ShouldResemble, []float64{0})

	fakegantry = &oneAxis{
		motor: &inject.Motor{
			PositionSupportedFunc: func(ctx context.Context) (bool, error) { return false, errors.New("not supported") },
			PositionFunc:          func(ctx context.Context) (float64, error) { return 1, errors.New("not supported") },
		},
		board:           createFakeBoard(),
		limitHigh:       true,
		limitSwitchPins: []string{"1", "2"},
		limitType:       switchLimitTypetwoPin,
		positionLimits:  []float64{0, 1},
		logger:          logger,
	}
	positions, err = fakegantry.GetPosition(ctx)
	test.That(t, positions, test.ShouldResemble, []float64{})
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry = &oneAxis{
		motor: &inject.Motor{
			PositionSupportedFunc: func(ctx context.Context) (bool, error) { return false, nil },
			PositionFunc:          func(ctx context.Context) (float64, error) { return 1, nil },
		},
		board: &inject.Board{
			GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) { return true, errors.New("not supported") },
		},
		limitHigh:       true,
		limitSwitchPins: []string{"1"},
		positionLimits:  []float64{0, 1},
		logger:          logger,
	}
	_, err = fakegantry.GetPosition(ctx)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry = &oneAxis{
		motor: &inject.Motor{
			PositionSupportedFunc: func(ctx context.Context) (bool, error) { return false, nil },
			PositionFunc:          func(ctx context.Context) (float64, error) { return 1, nil },
		},
		board: &inject.Board{
			GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) {
				if pin == "1" {
					return true, nil
				}
				return true, errors.New("not supported")
			},
		},
		limitHigh:       true,
		limitType:       switchLimitTypetwoPin,
		limitSwitchPins: []string{"1", "2"},
		positionLimits:  []float64{0, 1},
		logger:          logger,
	}
	_, err = fakegantry.GetPosition(ctx)
	test.That(t, err, test.ShouldNotBeNil)
}

func TestGetLengths(t *testing.T) {
	fakegantry := &oneAxis{
		lengthMm: float64(1.0),
	}
	ctx := context.Background()
	fakelengths, err := fakegantry.GetLengths(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.ShouldHaveLength(t, fakelengths, test.ShouldEqual(float64(1.0)))
}

func TestMoveToPosition(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	fakegantry := &oneAxis{
		logger:    logger,
		board:     createFakeBoard(),
		motor:     createFakeMotor(),
		limitHigh: true,
	}
	pos := []float64{1, 2}
	err := fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry MoveToPosition needs 1 position, got: 2")

	pos = []float64{1}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry position out of range, got 1.00 max is 0.00")

	fakegantry.lengthMm = float64(4)
	fakegantry.positionLimits = []float64{0, 4}
	fakegantry.limitSwitchPins = []string{"1", "2"}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err, test.ShouldBeNil)

	fakegantry.lengthMm = float64(4)
	fakegantry.positionLimits = []float64{0.01, .01}
	fakegantry.limitSwitchPins = []string{"1", "2"}
	fakegantry.motor = &inject.Motor{StopFunc: func(ctx context.Context) error { return errors.New("err") }}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.board = &inject.Board{GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) { return false, errors.New("err") }}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.board = &inject.Board{GetGPIOFunc: func(ctx context.Context, pin string) (bool, error) { return false, nil }}
	fakegantry.motor = &inject.Motor{
		StopFunc: func(ctx context.Context) error { return nil },
		GoToFunc: func(ctx context.Context, rpm float64, rotations float64) error { return errors.New("err") },
	}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err, test.ShouldNotBeNil)

	fakegantry.motor = &inject.Motor{GoToFunc: func(ctx context.Context, rpm float64, rotations float64) error { return nil }}
	err = fakegantry.MoveToPosition(ctx, pos)
	test.That(t, err, test.ShouldBeNil)
}

func TestModelFrame(t *testing.T) {
	fakegantry := &oneAxis{
		name:     "test",
		lengthMm: 1.0,
		axis:     r3.Vector{X: 0, Y: 0, Z: 1},
		model:    nil,
	}

	m := fakegantry.ModelFrame()
	test.That(t, m, test.ShouldNotBeNil)
}

func TestCurrentInputs(t *testing.T) {
	logger := golog.NewTestLogger(t)
	ctx := context.Background()

	fakegantry := &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1", "2"},
		lengthMm:        float64(200),
		positionLimits:  []float64{0, 2},
	}

	input, err := fakegantry.CurrentInputs(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, input[0].Value, test.ShouldEqual, 100)

	fakegantry = &oneAxis{
		motor:           createFakeMotor(),
		board:           createFakeBoard(),
		limitHigh:       true,
		logger:          logger,
		rpm:             float64(300),
		limitSwitchPins: []string{"1"},
		lengthMm:        float64(200),
		positionLimits:  []float64{0, 2},
		limitType:       switchLimitTypeOnePin,
	}

	input, err = fakegantry.CurrentInputs(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, input[0].Value, test.ShouldEqual, 100)

	// out of bounds position
	fakegantry = &oneAxis{
		motor:          &inject.Motor{PositionFunc: func(ctx context.Context) (float64, error) { return 5, errors.New("nope") }},
		board:          createFakeBoard(),
		limitHigh:      false,
		logger:         logger,
		rpm:            float64(300),
		lengthMm:       float64(200),
		positionLimits: []float64{0, 0.5},
	}

	input, err = fakegantry.CurrentInputs(ctx)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, input, test.ShouldBeNil)
}

func TestGoToInputs(t *testing.T) {
	ctx := context.Background()
	inputs := []referenceframe.Input{}
	logger := golog.NewTestLogger(t)

	fakegantry := &oneAxis{
		board:           createFakeBoard(),
		limitSwitchPins: []string{"1", "2"},
		limitHigh:       true,
		motor:           createFakeMotor(),
		lengthMm:        1.0,
		reductionRatio:  0.1,
		rpm:             10,
		axis:            r3.Vector{},
		limitType:       "",
		positionLimits:  []float64{1, 2},
		model:           nil,
		logger:          logger,
	}
	err := fakegantry.GoToInputs(ctx, inputs)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry MoveToPosition needs 1 position, got: 0")

	inputs = []referenceframe.Input{{Value: 1.0}, {Value: 2.0}}
	err = fakegantry.GoToInputs(ctx, inputs)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry MoveToPosition needs 1 position, got: 2")

	inputs = []referenceframe.Input{{Value: -1.0}}
	err = fakegantry.GoToInputs(ctx, inputs)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry position out of range, got -1.00 max is 1.00")

	inputs = []referenceframe.Input{{Value: 4.0}}
	err = fakegantry.GoToInputs(ctx, inputs)
	test.That(t, err.Error(), test.ShouldEqual, "oneAxis gantry position out of range, got 4.00 max is 1.00")

	inputs = []referenceframe.Input{{Value: 1.0}}
	err = fakegantry.GoToInputs(ctx, inputs)
	test.That(t, err, test.ShouldBeNil)
}
