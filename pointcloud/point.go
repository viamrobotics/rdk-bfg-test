package pointcloud

import (
	"image/color"
)

type Vec3 struct {
	X, Y, Z int
}

type Point interface {
	Position() Vec3

	HasColor() bool
	RGB255() (uint8, uint8, uint8)
	Color() color.Color

	HasValue() bool
	Value() int
}

type BasicPoint struct {
	position Vec3

	hasColor bool
	c        color.NRGBA

	hasValue bool
	value    int
}

func (bp *BasicPoint) Position() Vec3 {
	return bp.position
}

func (bp *BasicPoint) SetColor(c color.NRGBA) *BasicPoint {
	bp.hasColor = true
	bp.c = c
	return bp
}

func (bp *BasicPoint) HasColor() bool {
	return bp.hasColor
}

func (bp *BasicPoint) RGB255() (uint8, uint8, uint8) {
	return bp.c.R, bp.c.G, bp.c.B
}

func (bp *BasicPoint) Color() color.Color {
	return &bp.c
}

func (bp *BasicPoint) SetValue(v int) *BasicPoint {
	bp.hasValue = true
	bp.value = v
	return bp
}

func (bp *BasicPoint) HasValue() bool {
	return bp.hasValue
}

func (bp *BasicPoint) Value() int {
	return bp.value
}

func NewBasicPoint(x, y, z int) *BasicPoint {
	return &BasicPoint{position: Vec3{x, y, z}}
}

func NewColoredPoint(x, y, z int, c color.NRGBA) *BasicPoint {
	return &BasicPoint{position: Vec3{x, y, z}, c: c, hasColor: true}
}

func NewValuePoint(x, y, z int, v int) *BasicPoint {
	return &BasicPoint{position: Vec3{x, y, z}, value: v, hasValue: true}
}
