// Package ultrasonic implements an ultrasonic sensor based of the yahboom ultrasonic sensor
package ultrasonic

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"
	rdkutils "go.viam.com/utils"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
)

var modelname = resource.NewDefaultModel("ultrasonic")

// AttrConfig is used for converting config attributes.
type AttrConfig struct {
	TriggerPin    string `json:"trigger_pin"`
	EchoInterrupt string `json:"echo_interrupt_pin"`
	Board         string `json:"board"`
	TimeoutMs     uint   `json:"timeout_ms,omitempty"`
}

// Validate ensures all parts of the config are valid.
func (config *AttrConfig) Validate(path string) ([]string, error) {
	var deps []string
	if len(config.Board) == 0 {
		return nil, rdkutils.NewConfigValidationFieldRequiredError(path, "board")
	}
	deps = append(deps, config.Board)
	if len(config.TriggerPin) == 0 {
		return nil, rdkutils.NewConfigValidationFieldRequiredError(path, "trigger pin")
	}
	if len(config.EchoInterrupt) == 0 {
		return nil, rdkutils.NewConfigValidationFieldRequiredError(path, "echo interrupt pin")
	}
	return deps, nil
}

func init() {
	registry.RegisterComponent(
		sensor.Subtype,
		modelname,
		registry.Component{Constructor: func(
			ctx context.Context,
			deps registry.Dependencies,
			config config.Component,
			logger golog.Logger,
		) (interface{}, error) {
			return newSensor(ctx, deps, config.Name, config.ConvertedAttributes.(*AttrConfig))
		}})

	config.RegisterComponentAttributeMapConverter(sensor.Subtype, modelname,
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf AttrConfig
			return config.TransformAttributeMapToStruct(&conf, attributes)
		}, &AttrConfig{})
}

func newSensor(ctx context.Context, deps registry.Dependencies, name string, config *AttrConfig) (sensor.Sensor, error) {
	golog.Global().Debug("building ultrasonic sensor")
	s := &Sensor{Name: name, config: config}
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	s.cancelCtx = cancelCtx
	s.cancelFunc = cancelFunc

	res, ok := deps[board.Named(config.Board)]
	if !ok {
		return nil, errors.Errorf("ultrasonic: board %q missing from dependencies", config.Board)
	}

	b, ok := res.(board.Board)
	if !ok {
		return nil, errors.Errorf("ultrasonic: cannot find board %q", config.Board)
	}
	i, ok := b.DigitalInterruptByName(config.EchoInterrupt)
	if !ok {
		return nil, errors.Errorf("ultrasonic: cannot grab digital interrupt %q", config.EchoInterrupt)
	}
	g, err := b.GPIOPinByName(config.TriggerPin)
	if err != nil {
		return nil, errors.Wrapf(err, "ultrasonic: cannot grab gpio %q", config.TriggerPin)
	}
	s.echoInterrupt = i
	s.triggerPin = g

	if config.TimeoutMs > 0 {
		s.timeoutMs = config.TimeoutMs
	} else {
		// default to 1 sec
		s.timeoutMs = 1000
	}

	s.ticksChan = make(chan board.Tick, 2)
	if err := s.triggerPin.Set(ctx, false, nil); err != nil {
		return nil, errors.Wrap(err, "ultrasonic: cannot set trigger pin to low")
	}

	return s, nil
}

// Sensor ultrasonic sensor.
type Sensor struct {
	mu            sync.Mutex
	Name          string
	config        *AttrConfig
	echoInterrupt board.DigitalInterrupt
	triggerPin    board.GPIOPin
	ticksChan     chan board.Tick
	timeoutMs     uint
	cancelCtx     context.Context
	cancelFunc    func()
	generic.Unimplemented
}

func (s *Sensor) namedError(err error) error {
	return errors.Wrapf(
		err, "Error in ultrasonic sensor with name %s: ", s.Name,
	)
}

// Readings returns the calculated distance.
func (s *Sensor) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.echoInterrupt.AddCallback(s.ticksChan)
	defer s.echoInterrupt.RemoveCallback(s.ticksChan)

	// we send a high and a low to the trigger pin 10 microseconds
	// apart to signal the sensor to begin sending the sonic pulse
	if err := s.triggerPin.Set(ctx, true, nil); err != nil {
		return nil, s.namedError(errors.Wrap(err, "ultrasonic cannot set trigger pin to high"))
	}
	rdkutils.SelectContextOrWait(ctx, time.Microsecond*10)
	if err := s.triggerPin.Set(ctx, false, nil); err != nil {
		return nil, s.namedError(errors.Wrap(err, "ultrasonic cannot set trigger pin to low"))
	}
	// the first signal from the interrupt indicates that the sonic
	// pulse has been sent and the second indicates that the echo has been received
	var tick board.Tick
	ticks := make([]board.Tick, 2)
	for i := 0; i < 2; i++ {
		var signalStr string
		if i == 0 {
			signalStr = "sound pulse was emitted"
		} else {
			signalStr = "echo was received"
		}
		select {
		case tick = <-s.ticksChan:
			ticks[i] = tick
		case <-s.cancelCtx.Done():
			return nil, s.namedError(errors.New("ultrasonic: context canceled"))
		case <-time.After(time.Millisecond * time.Duration(s.timeoutMs)):
			return nil, s.namedError(errors.Errorf("timed out waiting for signal that %s", signalStr))
		}
	}
	timeB := ticks[0].TimestampNanosec
	timeA := ticks[1].TimestampNanosec
	// we calculate the distance to the nearest object based
	// on the time interval between the sound and its echo
	// and the speed of sound (343 m/s)
	secondsElapsed := float64(timeA-timeB) / math.Pow10(9)
	distMeters := secondsElapsed * 343.0 / 2.0
	return map[string]interface{}{"distance": distMeters}, nil
}

// Close remove interrupt callback of ultrasonic sensor.
func (s *Sensor) Close() error {
	s.cancelFunc()
	return nil
}
