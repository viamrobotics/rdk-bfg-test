package gpiostepper

import (
	"context"
	"sync"
	"testing"

	"github.com/edaniels/golog"
	"go.viam.com/test"
	"go.viam.com/utils/testutils"

	fakeboard "go.viam.com/rdk/components/board/fake"
	"go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/config"
)

func Test1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	logger := golog.NewTestLogger(t)

	b := &fakeboard.Board{GPIOPins: make(map[string]*fakeboard.GPIOPin)}

	mc := AttrConfig{}
	c := config.Component{
		Name: "fake_gpiostepper",
	}

	// Create motor with no board and default config
	t.Run("gpiostepper initializing test with no board and default config", func(t *testing.T) {
		_, err := newGPIOStepper(ctx, nil, mc, c.Name, logger)
		test.That(t, err, test.ShouldNotBeNil)
	})

	// Create motor with board and default config
	t.Run("gpiostepper initializing test with board and default config", func(t *testing.T) {
		_, err := newGPIOStepper(ctx, b, mc, c.Name, logger)
		test.That(t, err, test.ShouldNotBeNil)
	})

	mc.Pins = PinConfig{Direction: "b"}

	_, err := newGPIOStepper(ctx, b, mc, c.Name, logger)
	test.That(t, err, test.ShouldNotBeNil)

	mc.Pins.Step = "c"

	_, err = newGPIOStepper(ctx, b, mc, c.Name, logger)
	test.That(t, err, test.ShouldNotBeNil)

	mc.TicksPerRotation = 200

	mm, err := newGPIOStepper(ctx, b, mc, c.Name, logger)
	test.That(t, err, test.ShouldBeNil)

	m := mm.(*gpioStepper)

	t.Run("motor test supports position reporting", func(t *testing.T) {
		features, err := m.Properties(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, features[motor.PositionReporting], test.ShouldBeTrue)
	})

	t.Run("motor test isOn functionality", func(t *testing.T) {
		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, false)
		test.That(t, powerPct, test.ShouldEqual, 0.0)
	})

	t.Run("motor testing with positive rpm and positive revolutions", func(t *testing.T) {
		err = m.goForInternal(ctx, 100, 2)
		test.That(t, err, test.ShouldBeNil)

		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, true)
		test.That(t, powerPct, test.ShouldEqual, 1.0)

		testutils.WaitForAssertion(t, func(tb testing.TB) {
			tb.Helper()
			on, powerPct, err = m.IsPowered(ctx, nil)
			test.That(tb, err, test.ShouldBeNil)
			test.That(tb, on, test.ShouldEqual, false)
			test.That(tb, powerPct, test.ShouldEqual, 0.0)
		})

		pos, err := m.Position(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, pos, test.ShouldEqual, 2)
	})

	t.Run("motor testing with negative rpm and positive revolutions", func(t *testing.T) {
		err = m.goForInternal(ctx, -100, 2)
		test.That(t, err, test.ShouldBeNil)

		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, true)
		test.That(t, powerPct, test.ShouldEqual, 1.0)

		testutils.WaitForAssertion(t, func(tb testing.TB) {
			tb.Helper()
			on, powerPct, err = m.IsPowered(ctx, nil)
			test.That(tb, err, test.ShouldBeNil)
			test.That(tb, on, test.ShouldEqual, false)
			test.That(tb, powerPct, test.ShouldEqual, 0.0)
		})

		pos, err := m.Position(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, pos, test.ShouldEqual, 0)
	})

	t.Run("motor testing with positive rpm and negative revolutions", func(t *testing.T) {
		err = m.goForInternal(ctx, 100, -2)
		test.That(t, err, test.ShouldBeNil)

		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, true)
		test.That(t, powerPct, test.ShouldEqual, 1.0)

		testutils.WaitForAssertion(t, func(tb testing.TB) {
			tb.Helper()
			on, powerPct, err = m.IsPowered(ctx, nil)
			test.That(tb, err, test.ShouldBeNil)
			test.That(tb, on, test.ShouldEqual, false)
			test.That(tb, powerPct, test.ShouldEqual, 0.0)
		})

		pos, err := m.Position(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, pos, test.ShouldEqual, -2)
	})

	t.Run("motor testing with negative rpm and negative revolutions", func(t *testing.T) {
		err = m.goForInternal(ctx, -100, -2)
		test.That(t, err, test.ShouldBeNil)

		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, true)
		test.That(t, powerPct, test.ShouldEqual, 1.0)

		testutils.WaitForAssertion(t, func(tb testing.TB) {
			tb.Helper()
			on, powerPct, err = m.IsPowered(ctx, nil)
			test.That(tb, err, test.ShouldBeNil)
			test.That(tb, on, test.ShouldEqual, false)
			test.That(tb, powerPct, test.ShouldEqual, 0.0)
		})

		pos, err := m.Position(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, pos, test.ShouldEqual, 0)
	})
	t.Run("Ensure stop called when gofor is interrupted", func(t *testing.T) {
		ctx := context.Background()
		var wg sync.WaitGroup
		ctx, cancel := context.WithCancel(ctx)
		wg.Add(1)
		go func() {
			m.GoFor(ctx, 100, 100, map[string]interface{}{})
			wg.Done()
		}()
		cancel()
		wg.Wait()

		test.That(t, ctx.Err(), test.ShouldNotBeNil)
	})

	t.Run("motor testing with large # of revolutions", func(t *testing.T) {
		err = m.goForInternal(ctx, 100, 200)
		test.That(t, err, test.ShouldBeNil)

		on, powerPct, err := m.IsPowered(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, on, test.ShouldEqual, true)
		test.That(t, powerPct, test.ShouldEqual, 1.0)

		testutils.WaitForAssertion(t, func(tb testing.TB) {
			tb.Helper()
			pos, err := m.Position(ctx, nil)
			test.That(tb, err, test.ShouldBeNil)
			test.That(tb, pos, test.ShouldBeGreaterThan, 2)
		})

		err = m.Stop(ctx, nil)
		test.That(t, err, test.ShouldBeNil)

		pos, err := m.Position(ctx, nil)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, pos, test.ShouldBeGreaterThan, 2)
		test.That(t, pos, test.ShouldBeLessThan, 202)
	})

	t.Run("motor testing with 0 rpm", func(t *testing.T) {
		err = m.GoFor(ctx, 0, 1, nil)
		test.That(t, err, test.ShouldBeError, motor.NewZeroRPMError())
	})

	cancel()
}
