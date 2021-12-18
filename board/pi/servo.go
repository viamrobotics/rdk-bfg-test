//go:build pi
// +build pi

package pi

// #include <stdlib.h>
// #include <pigpio.h>
// #include "pi.h"
// #cgo LDFLAGS: -lpigpio
import "C"

import (
	"context"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"

	"go.viam.com/core/component/servo"
	"go.viam.com/core/config"
	piutils "go.viam.com/core/lib/pi"
	"go.viam.com/core/registry"
	"go.viam.com/core/robot"
)

// init registers a pi servo based on pigpio.
func init() {
	registry.RegisterComponent(servo.Subtype, "pi", registry.Component{Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (interface{}, error) {
		if !config.Attributes.Has("pin") {
			return nil, errors.New("expected pin for servo")
		}

		pin := config.Attributes.String("pin")
		bcom, have := piutils.BroadcomPinFromHardwareLabel(pin)
		if !have {
			return nil, errors.Errorf("no hw mapping for %s", pin)
		}

		return &piPigpioServo{C.uint(bcom)}, nil
	}})
}

// piPigpioServo implements a servo.Servo using pigpio.
type piPigpioServo struct {
	pin C.uint
}

func (s *piPigpioServo) Move(ctx context.Context, angle uint8) error {
	val := 500 + (2000.0 * float64(angle) / 180.0)
	res := C.gpioServo(s.pin, C.uint(val))
	if res != 0 {
		return errors.Errorf("gpioServo failed with %d", res)
	}
	return nil
}

func (s *piPigpioServo) AngularOffset(ctx context.Context) (uint8, error) {
	res := C.gpioGetServoPulsewidth(s.pin)
	if res <= 0 {
		// this includes, errors, we'll ignore
		return 0, nil
	}
	return uint8(180 * (float64(res) - 500.0) / 2000), nil
}
