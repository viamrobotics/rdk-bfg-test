// Package gps defines the interfaces of a GPS device which provides lat/long
// measurements.
package gps

import (
	"context"
	"math"
	"sync"

	"github.com/edaniels/golog"
	geo "github.com/kellydunn/golang-geo"
	viamutils "go.viam.com/utils"
	"go.viam.com/utils/rpc"

	"go.viam.com/rdk/component/generic"
	"go.viam.com/rdk/component/sensor"
	pb "go.viam.com/rdk/proto/api/component/gps/v1"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rlog"
	"go.viam.com/rdk/robot"
	"go.viam.com/rdk/subtype"
	"go.viam.com/rdk/utils"
)

func init() {
	registry.RegisterResourceSubtype(Subtype, registry.ResourceSubtype{
		Reconfigurable: WrapWithReconfigurable,
		RegisterRPCService: func(ctx context.Context, rpcServer rpc.Server, subtypeSvc subtype.Service) error {
			return rpcServer.RegisterServiceServer(
				ctx,
				&pb.GPSService_ServiceDesc,
				NewServer(subtypeSvc),
				pb.RegisterGPSServiceHandlerFromEndpoint,
			)
		},
		RPCServiceDesc: &pb.GPSService_ServiceDesc,
		RPCClient: func(ctx context.Context, conn rpc.ClientConn, name string, logger golog.Logger) interface{} {
			return NewClientFromConn(ctx, conn, name, logger)
		},
	})
}

// SubtypeName is a constant that identifies the component resource subtype string "gps".
const SubtypeName = resource.SubtypeName("gps")

// Subtype is a constant that identifies the component resource subtype.
var Subtype = resource.NewSubtype(
	resource.ResourceNamespaceRDK,
	resource.ResourceTypeComponent,
	SubtypeName,
)

// Named is a helper for getting the named GPS's typed resource name.
func Named(name string) resource.Name {
	return resource.NameFromSubtype(Subtype, name)
}

// A GPS represents a GPS that can report lat/long measurements.
type GPS interface {
	ReadLocation(ctx context.Context) (*geo.Point, error) // The current latitude and longitude
	ReadAltitude(ctx context.Context) (float64, error)    // The current altitude in meters
	ReadSpeed(ctx context.Context) (float64, error)       // Current ground speed in mm per sec
	generic.Generic
}

// A LocalGPS represents a GPS that can report accuracy, satellites and valid measurements.
type LocalGPS interface {
	GPS
	ReadAccuracy(ctx context.Context) (float64, float64, error) // Horizontal and vertical position error in meters
	ReadSatellites(ctx context.Context) (int, int, error)       // Number of satellites used for fix, and total in view
	ReadValid(ctx context.Context) (bool, error)                // Whether or not the GPS chip had a valid fix for the most recent dataset
}

var (
	_ = GPS(&reconfigurableGPS{})
	_ = LocalGPS(&reconfigurableLocalGPS{})
	_ = sensor.Sensor(&reconfigurableGPS{})
	_ = sensor.Sensor(&reconfigurableLocalGPS{})
	_ = resource.Reconfigurable(&reconfigurableGPS{})
	_ = resource.Reconfigurable(&reconfigurableLocalGPS{})
)

// FromDependencies is a helper for getting the named gps from a collection of
// dependencies.
func FromDependencies(deps registry.Dependencies, name string) (GPS, error) {
	res, ok := deps[Named(name)]
	if !ok {
		return nil, utils.DependencyNotFoundError(name)
	}
	part, ok := res.(GPS)
	if !ok {
		return nil, utils.DependencyTypeError(name, "GPS", res)
	}
	return part, nil
}

// FromRobot is a helper for getting the named GPS from the given Robot.
func FromRobot(r robot.Robot, name string) (GPS, error) {
	res, err := r.ResourceByName(Named(name))
	if err != nil {
		return nil, err
	}
	part, ok := res.(GPS)
	if !ok {
		return nil, utils.NewUnimplementedInterfaceError("GPS", res)
	}
	return part, nil
}

// NamesFromRobot is a helper for getting all GPS names from the given Robot.
func NamesFromRobot(r robot.Robot) []string {
	return robot.NamesBySubtype(r, Subtype)
}

// GetReadings is a helper for getting all readings from a GPS.
func GetReadings(ctx context.Context, g GPS) ([]interface{}, error) {
	loc, err := g.ReadLocation(ctx)
	if err != nil {
		return nil, err
	}
	alt, err := g.ReadAltitude(ctx)
	if err != nil {
		return nil, err
	}
	speed, err := g.ReadSpeed(ctx)
	if err != nil {
		return nil, err
	}

	readings := []interface{}{}
	if loc == nil {
		readings = append(readings, 0, 0)
	} else {
		readings = append(readings, loc.Lat(), loc.Lng())
	}

	readings = append(readings, alt, speed)

	localG, ok := g.(LocalGPS)
	if !ok {
		return readings, nil
	}

	active, total, err := localG.ReadSatellites(ctx)
	if err != nil {
		return nil, err
	}
	hAcc, vAcc, err := localG.ReadAccuracy(ctx)
	if err != nil {
		return nil, err
	}
	valid, err := localG.ReadValid(ctx)
	if err != nil {
		return nil, err
	}

	return append(readings, active, total, hAcc, vAcc, valid), nil
}

// GetHeading calculates bearing and absolute heading angles given 2 GPS coordinates
// 0 degrees indicate North, 90 degrees indicate East and so on.
func GetHeading(gps1 *geo.Point, gps2 *geo.Point, yawOffset float64) (float64, float64, float64) {
	// convert latitude and longitude readings from degrees to radians
	gps1Lat := utils.DegToRad(gps1.Lat())
	gps1Long := utils.DegToRad(gps1.Lng())
	gps2Lat := utils.DegToRad(gps2.Lat())
	gps2Long := utils.DegToRad(gps2.Lng())

	// calculate bearing from gps1 to gps 2
	dLon := gps2Long - gps1Long
	y := math.Sin(dLon) * math.Cos(gps2Lat)
	x := math.Cos(gps1Lat)*math.Sin(gps2Lat) - math.Sin(gps1Lat)*math.Cos(gps2Lat)*math.Cos(dLon)
	brng := utils.RadToDeg(math.Atan2(y, x))

	// maps bearing to 0-360 degrees
	if brng < 0 {
		brng += 360
	}

	// calculate absolute heading from bearing, accounting for yaw offset
	// e.g if the GPS antennas are mounted on the left and right sides of the robot,
	// the yaw offset would be roughly 90 degrees
	var standardBearing float64
	if brng > 180 {
		standardBearing = -(360 - brng)
	} else {
		standardBearing = brng
	}
	heading := brng - yawOffset

	// make heading positive again
	if heading < 0 {
		diff := math.Abs(heading)
		heading = 360 - diff
	}

	return brng, heading, standardBearing
}

type reconfigurableGPS struct {
	mu     sync.RWMutex
	actual GPS
}

func (r *reconfigurableGPS) Close(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return viamutils.TryClose(ctx, r.actual)
}

func (r *reconfigurableGPS) ProxyFor() interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual
}

func (r *reconfigurableGPS) Do(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.Do(ctx, cmd)
}

func (r *reconfigurableGPS) ReadLocation(ctx context.Context) (*geo.Point, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadLocation(ctx)
}

func (r *reconfigurableGPS) ReadAltitude(ctx context.Context) (float64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadAltitude(ctx)
}

func (r *reconfigurableGPS) ReadSpeed(ctx context.Context) (float64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadSpeed(ctx)
}

// GetReadings will use the default GPS GetReadings if not provided.
func (r *reconfigurableGPS) GetReadings(ctx context.Context) ([]interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if sensor, ok := r.actual.(sensor.Sensor); ok {
		return sensor.GetReadings(ctx)
	}
	return GetReadings(ctx, r.actual)
}

func (r *reconfigurableGPS) Reconfigure(ctx context.Context, newGPS resource.Reconfigurable) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.reconfigure(ctx, newGPS)
}

func (r *reconfigurableGPS) reconfigure(ctx context.Context, newGPS resource.Reconfigurable) error {
	actual, ok := newGPS.(*reconfigurableGPS)
	if !ok {
		return utils.NewUnexpectedTypeError(r, newGPS)
	}
	if err := viamutils.TryClose(ctx, r.actual); err != nil {
		rlog.Logger.Errorw("error closing old", "error", err)
	}
	r.actual = actual.actual
	return nil
}

type reconfigurableLocalGPS struct {
	*reconfigurableGPS
	actual LocalGPS
}

func (r *reconfigurableLocalGPS) Reconfigure(ctx context.Context, newGPS resource.Reconfigurable) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	gps, ok := newGPS.(*reconfigurableLocalGPS)
	if !ok {
		return utils.NewUnexpectedTypeError(r, newGPS)
	}
	if err := viamutils.TryClose(ctx, r.actual); err != nil {
		rlog.Logger.Errorw("error closing old", "error", err)
	}

	r.actual = gps.actual
	return r.reconfigurableGPS.reconfigure(ctx, gps.reconfigurableGPS)
}

func (r *reconfigurableLocalGPS) ReadSatellites(ctx context.Context) (int, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadSatellites(ctx)
}

func (r *reconfigurableLocalGPS) ReadAccuracy(ctx context.Context) (float64, float64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadAccuracy(ctx)
}

func (r *reconfigurableLocalGPS) ReadValid(ctx context.Context) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.actual.ReadValid(ctx)
}

// WrapWithReconfigurable converts a GPS to a reconfigurableGPS
// and a LocalGPS implementation to a reconfigurableLocalGPS.
// If GPS or LocalGPS is already a reconfigurableGPS, then nothing is done.
func WrapWithReconfigurable(r interface{}) (resource.Reconfigurable, error) {
	gps, ok := r.(GPS)
	if !ok {
		return nil, utils.NewUnimplementedInterfaceError("GPS", r)
	}
	if reconfigurable, ok := gps.(*reconfigurableGPS); ok {
		return reconfigurable, nil
	}
	rGPS := &reconfigurableGPS{actual: gps}
	gpsLocal, ok := r.(LocalGPS)
	if !ok {
		return rGPS, nil
	}
	if reconfigurable, ok := gps.(*reconfigurableLocalGPS); ok {
		return reconfigurable, nil
	}
	return &reconfigurableLocalGPS{actual: gpsLocal, reconfigurableGPS: rGPS}, nil
}
