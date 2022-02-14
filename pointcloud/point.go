package pointcloud

import (
	"image/color"
	"math"

	"github.com/golang/geo/r3"
	"gonum.org/v1/gonum/spatial/kdtree"
)

// Vec3 is a three-dimensional vector.
type Vec3 r3.Vector

// Vec3s is a series of three-dimensional vectors.
type Vec3s []Vec3

// Len returns the number of vectors.
func (vs Vec3s) Len() int {
	return len(vs)
}

// Swap swaps two vectors positionally.
func (vs Vec3s) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

// Less returns which vector is less than the other based on
// r3.Vector.Cmp.
func (vs Vec3s) Less(i, j int) bool {
	cmp := (r3.Vector)(vs[i]).Cmp((r3.Vector)(vs[j]))
	if cmp == 0 {
		return false
	}
	return cmp < 0
}

// A Point describes a single point within a PointCloud. It is the
// collection of these points that forms the cloud.
type Point interface {
	// Position is the vector describing where the point is in the cloud.
	Position() Vec3

	// Clone copies the point to a new position.
	Clone(v Vec3) Point

	// HasColor returns whether or not this point is colored.
	HasColor() bool

	// RGB255 returns, if colored, the RGB components of the color. There
	// is no alpha channel right now and as such the data can be assumed to be
	// premultiplied.
	RGB255() (uint8, uint8, uint8)

	// Color returns the native color of the point.
	Color() color.Color

	// SetColor sets the given color on the point.
	// Note(erd): we should try to remove this in favor of immutability.
	SetColor(c color.NRGBA) Point

	// HasValue returns whether or not this point has some user data value
	// associated with it.
	HasValue() bool

	// Value returns the user data set value, if it exists.
	Value() int

	// SetValue sets the given user data value on the point.
	// Note(erd): we should try to remove this in favor of immutability.
	SetValue(v int) Point

	// Value returns the intesity value, or 0 if it doesn't exist
	Intensity() uint16

	// SetIntensity sets the intensity on the point.
	SetIntensity(v uint16) Point

	// Distance returns the distance between the the current point and the given point.
	// Used to fulfill kdtree.Comparable.
	Distance(p kdtree.Comparable) float64

	// Dims returns the dimensions of the point. Used to fulfill kdtree.Comparable.
	Dims() int

	// Compare returns the signed distance from the plane passing through p and perpendicular to dim d.
	// Used to fulfill kdtree.Comparable.
	Compare(p kdtree.Comparable, d kdtree.Dim) float64
}

type basicPoint struct {
	position Vec3

	hasColor bool
	c        color.NRGBA

	hasValue bool
	value    int

	intensity uint16
}

// NewBasicPoint returns a point that is solely positionally based.
func NewBasicPoint(x, y, z float64) Point {
	return &basicPoint{position: Vec3{x, y, z}}
}

// NewColoredPoint returns a point that has both position and color.
func NewColoredPoint(x, y, z float64, c color.NRGBA) Point {
	return &basicPoint{position: Vec3{x, y, z}, c: c, hasColor: true}
}

// NewValuePoint returns a point that has both position and a user data value.
func NewValuePoint(x, y, z float64, v int) Point {
	return &basicPoint{position: Vec3{x, y, z}, value: v, hasValue: true}
}

// Clone copies the point over with a new position.
func (bp *basicPoint) Clone(v Vec3) Point {
	return &basicPoint{
		position:  v,
		hasColor:  bp.hasColor,
		c:         bp.c,
		hasValue:  bp.hasValue,
		value:     bp.value,
		intensity: bp.intensity,
	}
}

func (bp *basicPoint) Position() Vec3 {
	return bp.position
}

// Distance returns the distance between the the current point and the given point.
// Used to fulfill kdtree.Comparable.
func (bp *basicPoint) Distance(p kdtree.Comparable) float64 {
	pp, ok := p.(Point)
	if !ok {
		panic("kdtree.Comparable fed into basicPoint.Distance is not a Point")
	}
	v1, v2 := bp.Position(), pp.Position()
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2) + math.Pow(v2.Z-v1.Z, 2))
}

// Dims returns the dimensions of the point. Used to fulfill kdtree.Comparable.
func (bp *basicPoint) Dims() int { return 3 }

// Compare returns the signed distance from the plane passing through p and perpendicular to dim d.
// Used to fulfill kdtree.Comparable.
func (bp *basicPoint) Compare(p kdtree.Comparable, d kdtree.Dim) float64 {
	p2, ok := p.(Point)
	if !ok {
		panic("kdtree.Comparable fed into basicPoint.Distance is not a Point")
	}
	v1, v2 := bp.Position(), p2.Position()
	switch d {
	case 0:
		return v1.X - v2.X
	case 1:
		return v1.Y - v2.Y
	case 2:
		return v1.Z - v2.Z
	default:
		panic("illegal dimension fed to basicPoint.Compare")
	}
}

func (bp *basicPoint) SetColor(c color.NRGBA) Point {
	bp.hasColor = true
	bp.c = c
	return bp
}

func (bp *basicPoint) HasColor() bool {
	return bp.hasColor
}

func (bp *basicPoint) RGB255() (uint8, uint8, uint8) {
	return bp.c.R, bp.c.G, bp.c.B
}

func (bp *basicPoint) Color() color.Color {
	return &bp.c
}

func (bp *basicPoint) SetValue(v int) Point {
	bp.hasValue = true
	bp.value = v
	return bp
}

func (bp *basicPoint) HasValue() bool {
	return bp.hasValue
}

func (bp *basicPoint) Value() int {
	return bp.value
}

func (bp *basicPoint) SetIntensity(v uint16) Point {
	bp.intensity = v
	return bp
}

func (bp *basicPoint) Intensity() uint16 {
	return bp.intensity
}

// GetPositions gets the positions of the slice of points.
func GetPositions(pts []Point) []r3.Vector {
	positions := make([]r3.Vector, len(pts))
	for i, pt := range pts {
		positions[i] = r3.Vector(pt.Position())
	}
	return positions
}
