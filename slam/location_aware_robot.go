package slam

import (
	"fmt"
	"image"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/viamrobotics/robotcore/base"
	"github.com/viamrobotics/robotcore/lidar"
	"github.com/viamrobotics/robotcore/robots/fake"
	"github.com/viamrobotics/robotcore/utils"

	"github.com/edaniels/golog"
)

// relative to first device
type DeviceOffset struct {
	Angle                float64
	DistanceX, DistanceY float64
}

type LocationAwareRobot struct {
	base            base.Base
	baseOrientation int // relative to map
	basePosX        int
	basePosY        int
	maxBounds       image.Point

	devices       []lidar.Device
	deviceOffsets []DeviceOffset
	orientations  []float64

	area          *SquareArea
	areaBounds    image.Point
	distinctAreas []*SquareArea

	clientDeviceNum int
	clientZoom      float64
	clientClickMode string

	isMoving         int32
	moveMu           sync.Mutex
	scanMu           sync.Mutex
	clientSettingsMu sync.Mutex
	closeCh          chan struct{}
	closeOnce        sync.Once
}

func NewLocationAwareRobot(
	base base.Base,
	baseStart image.Point,
	devices []lidar.Device,
	deviceOffsets []DeviceOffset,
	area *SquareArea,
	areaBounds image.Point,
) (*LocationAwareRobot, error) {
	areaSize, areaSizeScale := area.Size()
	distinctAreas := make([]*SquareArea, 0, len(devices))
	for range devices {
		distinctAreas = append(distinctAreas, NewSquareArea(areaSize, areaSizeScale))
	}

	var maxBoundsX, maxBoundsY int
	for _, dev := range devices {
		bounds, err := dev.Bounds()
		if err != nil {
			return nil, err
		}
		if bounds.X > maxBoundsX {
			maxBoundsX = bounds.X
		}
		if bounds.Y > maxBoundsY {
			maxBoundsY = bounds.Y
		}
	}

	return &LocationAwareRobot{
		base:      base,
		basePosX:  baseStart.X,
		basePosY:  baseStart.Y,
		maxBounds: image.Point{maxBoundsX, maxBoundsY},

		devices:       devices,
		deviceOffsets: deviceOffsets,
		orientations:  make([]float64, len(devices)),

		area:          area,
		areaBounds:    areaBounds,
		distinctAreas: distinctAreas,

		clientDeviceNum: -1,
		clientZoom:      1,
		clientClickMode: clientClickModeInfo,
		closeCh:         make(chan struct{}),
	}, nil
}

func (lar *LocationAwareRobot) Start() {
	select {
	case <-lar.closeCh:
		return
	default:
	}
	lar.cullLoop()
	lar.updateLoop()
}

func (lar *LocationAwareRobot) Stop() {
	lar.closeOnce.Do(func() {
		close(lar.closeCh)
	})
}

func (lar *LocationAwareRobot) Close() error {
	lar.Stop()
	return nil
}

func (lar *LocationAwareRobot) Move(amount *int, rotateTo *Direction) error {
	lar.scanMu.Lock()
	atomic.StoreInt32(&lar.isMoving, 1)
	lar.scanMu.Unlock()
	defer atomic.StoreInt32(&lar.isMoving, 0)
	lar.moveMu.Lock()
	defer lar.moveMu.Unlock()

	move := base.Move{Speed: 0, Block: true}

	if rotateTo != nil {
		golog.Global.Debugw("request to rotate", "dir", *rotateTo)
		from := lar.baseOrientation
		var to int
		switch *rotateTo {
		case DirectionUp:
			to = 0
		case DirectionRight:
			to = 90
		case DirectionDown:
			to = 180
		case DirectionLeft:
			to = 270
		default:
			return fmt.Errorf("do not know how to rotate to absolute %q", *rotateTo)
		}
		rotateBy := from - to
		if rotateBy != 180 && rotateBy != -180 {
			rotateBy = (rotateBy + 180) % 180
			if from > to {
				rotateBy *= -1
			}
		}
		move.AngleDeg = rotateBy
	}
	newOrientation := (((lar.baseOrientation + move.AngleDeg) % 360) + 360) % 360

	newX := lar.basePosX
	newY := lar.basePosY
	if amount != nil {
		actualAmount := *amount
		errMsg := fmt.Errorf("cannot move at orientation %d; stuck", newOrientation)
		switch newOrientation {
		case 0:
			if lar.basePosY-actualAmount < 0 {
				return errMsg
			}
			golog.Global.Debugw("up", "amount", actualAmount)
			newY = lar.basePosY - actualAmount
		case 90:
			if lar.basePosX+actualAmount >= lar.areaBounds.X {
				return errMsg
			}
			golog.Global.Debugw("right", "amount", actualAmount)
			newX = lar.basePosX + actualAmount
		case 180:
			if lar.basePosY+actualAmount >= lar.areaBounds.Y {
				return errMsg
			}
			golog.Global.Debugw("down", "amount", actualAmount)
			newY = lar.basePosY + actualAmount
		case 270:
			if lar.basePosX-actualAmount < 0 {
				return errMsg
			}
			golog.Global.Debugw("left", "amount", actualAmount)
			newX = lar.basePosX - actualAmount
		default:
			return fmt.Errorf("cannot move at orientation %d", newOrientation)
		}
		move.DistanceMM = actualAmount * 10
	}

	if newX != lar.basePosX || newY != lar.basePosY {
		// TODO(erd): refactor out to func
		// detect obstacle START

		// TODO(erd): use area of entity to determine collision
		// the lidar will give out around this distance so
		// we must make sure to not approach an area like this so as
		// to avoid the collision disappearing.
		const detectionBuffer = 15

		// Version 1: detect if single point passes through any straight line
		// Note: straight only works because we use 90 deg rotations
		var collides bool
		lar.area.Mutate(func(mutArea MutableArea) {
			mutArea.DoNonZero(func(x, y int, v float64) {
				switch newOrientation {
				case 0:
					if newX == x && y < lar.basePosY && y > int(math.Max(0, float64(newY-detectionBuffer))) {
						collides = true
					}
				case 90:
					if newY == y && x > lar.basePosX && x < int(math.Min(math.MaxFloat64, float64(newX+detectionBuffer))) {
						collides = true
					}
				case 180:
					if newX == x && y > lar.basePosY && y < int(math.Min(0, float64(newY+detectionBuffer))) {
						collides = true
					}
				case 270:
					if newY == y && x < lar.basePosX && x > int(math.Max(math.MaxFloat64, float64(newX+detectionBuffer))) {
						collides = true
					}
				}
			})
		})
		if collides {
			return fmt.Errorf("cannot move to (%d,%d) via %d; would collide", newX, newY, newOrientation)
		}

		// detect obstacle END
	}

	if _, _, err := base.DoMove(move, lar.base); err != nil {
		return err
	}
	lar.basePosX = newX
	lar.basePosY = newY
	lar.baseOrientation = newOrientation
	return nil
}

func (lar *LocationAwareRobot) basePos() (int, int) {
	return lar.basePosX, lar.basePosY
}

func (lar *LocationAwareRobot) String() string {
	return fmt.Sprintf("pos: (%d, %d)", lar.basePosX, lar.basePosY)
}

func (lar *LocationAwareRobot) cullLoop() {
	_, scaleDown := lar.area.Size()
	maxBoundsX := lar.maxBounds.X * scaleDown
	maxBoundsY := lar.maxBounds.Y * scaleDown

	cull := func() {
		if atomic.LoadInt32(&lar.isMoving) == 1 {
			return
		}
		// TODO(erd): not thread safe
		basePosX, basePosY := lar.basePos()

		// calculate ideal visibility bounds
		areaMinX := basePosX - maxBoundsX/2
		areaMaxX := basePosX + maxBoundsX/2
		areaMinY := basePosY - maxBoundsY/2
		areaMaxY := basePosY + maxBoundsY/2

		// decrement observable area which will be refreshed by scans
		// within the area (assuming the lidar is active)
		cullArea := func(area *SquareArea, minX, maxX, minY, maxY int) {
			area.Mutate(func(mutArea MutableArea) {
				mutArea.DoNonZero(func(x, y int, v float64) {
					if x < minX || x > maxX || y < minY || y > maxY {
						return
					}
					mutArea.Set(x, y, v-1)
				})
			})
		}

		cullArea(lar.area, areaMinX, areaMaxX, areaMinY, areaMaxY)

		for i, area := range lar.distinctAreas {
			bounds, err := lar.devices[i].Bounds()
			if err != nil {
				panic(err)
			}
			bounds.X *= scaleDown
			bounds.Y *= scaleDown

			areaMinX := basePosX - bounds.X/2
			areaMaxX := basePosX + bounds.X/2
			areaMinY := basePosY - bounds.Y/2
			areaMaxY := basePosY + bounds.Y/2

			cullArea(area, areaMinX, areaMaxX, areaMinY, areaMaxY)
		}
	}

	ticker := time.NewTicker(time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-lar.closeCh:
				return
			default:
			}
			select {
			case <-lar.closeCh:
				return
			case <-ticker.C:
			}
			cull()
		}
	}()
}

const cullTTL = 3

func (lar *LocationAwareRobot) updateLoop() {
	update := func() {
		if atomic.LoadInt32(&lar.isMoving) == 1 {
			return
		}
		basePosX, basePosY := lar.basePos()

		for _, dev := range lar.devices {
			if fake, ok := dev.(*fake.Lidar); ok {
				fake.SetPosition(image.Point{basePosX, basePosY})
			}
		}
		allMeasurements := make([]lidar.Measurements, len(lar.devices))
		for i, dev := range lar.devices {
			measurements, err := dev.Scan()
			if err != nil {
				golog.Global.Debugw("bad scan", "device", i, "error", err)
				continue
			}
			allMeasurements[i] = measurements
		}

		areaSize, scaleDown := lar.area.Size()
		areaSize *= scaleDown
		for i, measurements := range allMeasurements {
			minAngle := math.Inf(1)
			var adjust bool
			var offsets DeviceOffset
			if i != 0 && i-1 < len(lar.deviceOffsets) {
				offsets = lar.deviceOffsets[i-1]
				adjust = true
			}
			// TODO(erd): better to just adjust in advance?
			for _, next := range measurements {
				angle := next.Angle()
				x, y := next.Coords()
				if adjust {
					angle += offsets.Angle
					angleRad := utils.DegToRad(offsets.Angle)
					// rotate vector around base ccw
					newX := math.Cos(angleRad)*x - math.Sin(angleRad)*y
					newY := math.Sin(angleRad)*x + math.Cos(angleRad)*y
					x = newX
					y = newY
				}
				if angle < minAngle {
					minAngle = angle
				}
				detectedX := int(float64(basePosX) + offsets.DistanceX + x*float64(scaleDown))
				detectedY := int(float64(basePosY) + offsets.DistanceY + y*float64(scaleDown))
				if detectedX < 0 || detectedX >= areaSize {
					continue
				}
				if detectedY < 0 || detectedY >= areaSize {
					continue
				}
				// TODO(erd): should we also add here as a sense of permanency
				// Want to also combine this with occlusion, right. So if there's
				// a wall detected, and we're pretty confident it's staying there,
				// it being occluded should give it a low chance of it being removed.
				// Realistically once the bounds of a location are determined, most
				// environments would only have it deform over very long periods of time.
				// Probably longer than the lifetime of the application itself.
				lar.area.Mutate(func(area MutableArea) {
					area.Set(detectedX, detectedY, cullTTL)
				})
				lar.distinctAreas[i].Mutate(func(area MutableArea) {
					area.Set(detectedX, detectedY, cullTTL)
				})
			}
			lar.orientations[i] = minAngle
		}
	}
	ticker := time.NewTicker(33 * time.Millisecond)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-lar.closeCh:
				return
			default:
			}
			select {
			case <-lar.closeCh:
				return
			case <-ticker.C:
			}
			update()
		}
	}()
}

func (lar *LocationAwareRobot) areaToView() (image.Point, *SquareArea, error) {
	devNum := lar.getClientDeviceNum()
	if devNum == -1 {
		return lar.maxBounds, lar.area, nil
	}
	dev := lar.devices[devNum]
	bounds, err := dev.Bounds()
	if err != nil {
		return image.Point{}, nil, err
	}
	return bounds, lar.distinctAreas[devNum], nil
}

func (lar *LocationAwareRobot) getClientDeviceNum() int {
	lar.clientSettingsMu.Lock()
	defer lar.clientSettingsMu.Unlock()
	return lar.clientDeviceNum
}

func (lar *LocationAwareRobot) setClientDeviceNumber(num int) {
	lar.clientSettingsMu.Lock()
	defer lar.clientSettingsMu.Unlock()
	lar.clientDeviceNum = num
}

func (lar *LocationAwareRobot) rotateTo(dir Direction) error {
	return lar.Move(nil, &dir)
}
