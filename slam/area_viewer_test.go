package slam

import (
	"context"
	"testing"

	"github.com/edaniels/test"
)

func TestAreaViewer(t *testing.T) {
	sa := NewSquareArea(100, 10)
	viewer := AreaViewer{sa}

	sa.Mutate(func(area MutableArea) {
		area.Set(50, 50, 1)
		area.Set(500, 90, 1)
		area.Set(222, 840, 1)
	})

	img, err := viewer.Next(context.Background())
	test.That(t, err, test.ShouldBeNil)
	bounds := img.Bounds()
	test.That(t, bounds.Dx(), test.ShouldEqual, 1000)
	test.That(t, bounds.Dy(), test.ShouldEqual, 1000)
	p1 := img.At(50, 50)
	p2 := img.At(500, 90)
	p3 := img.At(222, 840)
	r1, g1, b1, _ := p1.RGBA()
	r2, g2, b2, _ := p2.RGBA()
	r3, g3, b3, _ := p3.RGBA()

	test.That(t, r1>>8, test.ShouldEqual, 255)
	test.That(t, r2>>8, test.ShouldEqual, 255)
	test.That(t, r3>>8, test.ShouldEqual, 255)
	test.That(t, g1, test.ShouldEqual, 0)
	test.That(t, g2, test.ShouldEqual, 0)
	test.That(t, g3, test.ShouldEqual, 0)
	test.That(t, b1, test.ShouldEqual, 0)
	test.That(t, b2, test.ShouldEqual, 0)
	test.That(t, b3, test.ShouldEqual, 0)

	test.That(t, viewer.Close(), test.ShouldBeNil)
}
