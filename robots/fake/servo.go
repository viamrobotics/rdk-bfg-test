package fake

import (
	"context"

	"github.com/edaniels/golog"

	"go.viam.com/core/config"
	"go.viam.com/core/registry"
	"go.viam.com/core/robot"
	"go.viam.com/core/servo"
)

func init() {
	registry.RegisterServo(modelName, registry.Servo{Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (servo.Servo, error) {
		return &Servo{Name: config.Name}, nil
	}})
}

// A Servo allows setting and reading a single angle.
type Servo struct {
	Name  string
	angle uint8
}

// Move sets the given angle.
func (s *Servo) Move(ctx context.Context, angle uint8) error {
	s.angle = angle
	return nil
}

// Current returns the set angle.
func (s *Servo) Current(ctx context.Context) (uint8, error) {
	return s.angle, nil
}
