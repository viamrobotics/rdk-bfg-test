package nmea

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"go.viam.com/utils"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/spatialmath"
)

// I2CAttrConfig is used for converting Serial NMEA MovementSensor config attributes.
type I2CAttrConfig struct {
	// I2C
	Board   string `json:"board"`
	Bus     string `json:"bus"`
	I2cAddr int    `json:"i2c_addr"`
}

// ValidateI2C ensures all parts of the config are valid.
func (config *I2CAttrConfig) ValidateI2C(path string) error {
	if len(config.Board) == 0 {
		return errors.New("expected nonempty board")
	}
	if len(config.Bus) == 0 {
		return errors.New("expected nonempty bus")
	}
	if config.I2cAddr == 0 {
		return errors.New("expected nonempty i2c address")
	}

	return nil
}

func init() {
	registry.RegisterComponent(
		movementsensor.Subtype,
		"nmea-pmtkI2C",
		registry.Component{Constructor: func(
			ctx context.Context,
			deps registry.Dependencies,
			config config.Component,
			logger golog.Logger,
		) (interface{}, error) {
			return newPmtkI2CNMEAMovementSensor(ctx, deps, config, logger)
		}})
}

// PmtkI2CNMEAMovementSensor allows the use of any MovementSensor chip that communicates over I2C using the PMTK protocol.
type PmtkI2CNMEAMovementSensor struct {
	generic.Unimplemented
	mu          sync.RWMutex
	bus         board.I2C
	addr        byte
	wbaud       int
	logger      golog.Logger
	disableNmea bool

	data gpsData

	cancelCtx               context.Context
	cancelFunc              func()
	activeBackgroundWorkers sync.WaitGroup

	errMu     sync.Mutex
	lastError error
}

func newPmtkI2CNMEAMovementSensor(
	ctx context.Context,
	deps registry.Dependencies,
	config config.Component,
	logger golog.Logger,
) (nmeaMovementSensor, error) {
	b, err := board.FromDependencies(deps, config.Attributes.String("board"))
	if err != nil {
		return nil, fmt.Errorf("gps init: failed to find board: %w", err)
	}
	localB, ok := b.(board.LocalBoard)
	if !ok {
		return nil, fmt.Errorf("board %s is not local", config.Attributes.String("board"))
	}
	i2cbus, ok := localB.I2CByName(config.Attributes.String("bus"))
	if !ok {
		return nil, fmt.Errorf("gps init: failed to find i2c bus %s", config.Attributes.String("bus"))
	}
	addr := config.Attributes.Int("i2c_addr", -1)
	if addr == -1 {
		return nil, errors.New("must specify gps i2c address")
	}
	wbaud := config.Attributes.Int("ntrip_baud", 38400)
	disableNmea := config.Attributes.Bool(disableNmeaName, false)
	if disableNmea {
		logger.Info("SerialNMEAMovementSensor: NMEA reading disabled")
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	g := &PmtkI2CNMEAMovementSensor{
		bus:         i2cbus,
		addr:        byte(addr),
		wbaud:       wbaud,
		cancelCtx:   cancelCtx,
		cancelFunc:  cancelFunc,
		logger:      logger,
		disableNmea: disableNmea,
	}

	if err := g.Start(ctx); err != nil {
		return nil, err
	}
	return g, g.lastError
}

func (g *PmtkI2CNMEAMovementSensor) setLastError(err error) {
	g.errMu.Lock()
	defer g.errMu.Unlock()

	g.lastError = err
}

// Start begins reading nmea messages from module and updates gps data.
func (g *PmtkI2CNMEAMovementSensor) Start(ctx context.Context) error {
	handle, err := g.bus.OpenHandle(g.addr)
	if err != nil {
		g.logger.Errorf("can't open gps i2c %s", err)
		return err
	}
	// Send GLL, RMC, VTG, GGA, GSA, and GSV sentences each 1000ms
	baudcmd := fmt.Sprintf("PMTK251,%d", g.wbaud)
	cmd251 := addChk([]byte(baudcmd))
	cmd314 := addChk([]byte("PMTK314,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0"))
	cmd220 := addChk([]byte("PMTK220,1000"))

	err = handle.Write(ctx, cmd251)
	if err != nil {
		g.logger.Debug("Failed to set baud rate")
	}
	err = handle.Write(ctx, cmd314)
	if err != nil {
		g.logger.Errorf("i2c handle write failed %s", err)
		return err
	}
	err = handle.Write(ctx, cmd220)
	if err != nil {
		g.logger.Errorf("i2c handle write failed %s", err)
		return err
	}
	err = handle.Close()
	if err != nil {
		g.logger.Errorf("failed to close handle: %s", err)
		return err
	}

	g.activeBackgroundWorkers.Add(1)
	utils.PanicCapturingGo(func() {
		defer g.activeBackgroundWorkers.Done()
		strBuf := ""
		for {
			select {
			case <-g.cancelCtx.Done():
				return
			default:
			}

			if !g.disableNmea {
				// Opening an i2c handle blocks the whole bus, so we open/close each loop so other things also have a chance to use it
				handle, err := g.bus.OpenHandle(g.addr)
				if err != nil {
					g.logger.Errorf("can't open gps i2c handle: %s", err)
					g.setLastError(err)
					return
				}
				buffer, err := handle.Read(ctx, 1024)
				hErr := handle.Close()
				if hErr != nil {
					g.logger.Errorf("failed to close handle: %s", hErr)
					g.setLastError(err)
					return
				}
				if err != nil {
					g.logger.Error(err)
					continue
				}
				for _, b := range buffer {
					// PMTK uses CRLF line endings to terminate sentences, but just LF to blank data.
					// Since CR should never appear except at the end of our sentence, we use that to determine sentence end.
					// LF is merely ignored.
					if b == 0x0D {
						if strBuf != "" {
							g.mu.Lock()
							err = g.data.parseAndUpdate(strBuf)
							g.mu.Unlock()
							if err != nil {
								g.logger.Debugf("can't parse nmea : %s, %v", strBuf, err)
							}
						}
						strBuf = ""
					} else if b != 0x0A && b != 0xFF { // adds only valid bytes
						strBuf += string(b)
					}
				}
			}
		}
	})

	return g.lastError
}

// GetBusAddr returns the bus and address that takes in rtcm corrections.
func (g *PmtkI2CNMEAMovementSensor) GetBusAddr() (board.I2C, byte) {
	return g.bus, g.addr
}

// GetPosition returns the current geographic location of the MovementSensor.
func (g *PmtkI2CNMEAMovementSensor) GetPosition(ctx context.Context) (*geo.Point, float64, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.data.location, g.data.alt, g.lastError
}

// GetAccuracy returns the accuracy, hDOP and vDOP.
func (g *PmtkI2CNMEAMovementSensor) GetAccuracy(ctx context.Context) (map[string]float32, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return map[string]float32{"hDOP": float32(g.data.hDOP), "vDOP": float32(g.data.vDOP)}, g.lastError
}

// GetLinearVelocity returns the current speed of the MovementSensor.
func (g *PmtkI2CNMEAMovementSensor) GetLinearVelocity(ctx context.Context) (r3.Vector, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return r3.Vector{0, g.data.speed, 0}, g.lastError
}

// GetAngularVelocity not supported.
func (g *PmtkI2CNMEAMovementSensor) GetAngularVelocity(ctx context.Context) (spatialmath.AngularVelocity, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return spatialmath.AngularVelocity{}, g.lastError
}

// GetCompassHeading not supported.
func (g *PmtkI2CNMEAMovementSensor) GetCompassHeading(ctx context.Context) (float64, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return 0, g.lastError
}

// GetOrientation not supporter.
func (g *PmtkI2CNMEAMovementSensor) GetOrientation(ctx context.Context) (spatialmath.Orientation, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return nil, g.lastError
}

// GetProperties what can I do!
func (g *PmtkI2CNMEAMovementSensor) GetProperties(ctx context.Context) (*movementsensor.Properties, error) {
	return &movementsensor.Properties{
		LinearVelocitySupported: true,
		PositionSupported:       true,
	}, g.lastError
}

// ReadFix returns quality.
func (g *PmtkI2CNMEAMovementSensor) ReadFix(ctx context.Context) (int, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.data.fixQuality, g.lastError
}

// GetReadings will use return all of the MovementSensor Readings.
func (g *PmtkI2CNMEAMovementSensor) GetReadings(ctx context.Context) (map[string]interface{}, error) {
	readings, err := movementsensor.GetReadings(ctx, g)
	if err != nil {
		return nil, err
	}

	fix, err := g.ReadFix(ctx)
	if err != nil {
		return nil, err
	}

	readings["fix"] = fix

	return readings, nil
}

// Close shuts down the SerialNMEAMOVEMENTSENSOR.
func (g *PmtkI2CNMEAMovementSensor) Close() error {
	g.cancelFunc()
	g.activeBackgroundWorkers.Wait()

	return g.lastError
}

// PMTK checksums commands by XORing together each byte.
func addChk(data []byte) []byte {
	chk := checksum(data)
	newCmd := []byte("$")
	newCmd = append(newCmd, data...)
	newCmd = append(newCmd, []byte("*")...)
	newCmd = append(newCmd, chk)
	return newCmd
}

func checksum(data []byte) byte {
	var chk byte
	for _, b := range data {
		chk ^= b
	}
	return chk
}
