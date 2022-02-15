package motor

import "github.com/pkg/errors"

// NewGoTillStopUnsupportedError returns a standard error for when a motor
// is required to support the GoTillStop feature.
func NewGoTillStopUnsupportedError(motorName string) error {
	return errors.Errorf("motor with name %s does not support GoTillStop", motorName)
}

// NewResetZeroPositionUnsupportedError returns a standard error for when a motor
// is required to support reseting the zero position.
func NewResetZeroPositionUnsupportedError(motorName string) error {
	return errors.Errorf("motor with name %s does not support ResetZeroPosition", motorName)
}

// NewFeatureUnsupportedError returns an error representing the need
// for a motor to support a particular feature.
func NewFeatureUnsupportedError(feature Feature, motorName string) error {
	return errors.Errorf("motor named %s must support feature motor.%s", motorName, feature)
}
