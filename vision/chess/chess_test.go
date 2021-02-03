package chess

import (
	"image"
	"testing"

	"github.com/disintegration/imaging"

	"github.com/viamrobotics/robotcore/vision"
	"github.com/viamrobotics/robotcore/vision/segmentation"
)

type P func(vision.Image) (image.Image, []image.Point, error)

type ChessImageProcessDebug struct {
	p P
}

func (dd ChessImageProcessDebug) Process(d *vision.MultipleImageTestDebugger, fn string, img vision.Image) error {
	out, corners, err := dd.p(img)
	if err != nil {
		return err
	}

	d.GotDebugImage(out, "corners")

	if corners != nil {
		warped, _, err := warpColorAndDepthToChess(img, &vision.DepthMap{}, corners)
		if err != nil {
			return err
		}

		d.GotDebugImage(warped.Image(), "warped")

		starts := []image.Point{}
		for x := 50; x <= 750; x += 100 {
			for y := 50; y <= 750; y += 100 {
				starts = append(starts, image.Point{x, y})
			}
		}

		i2 := imaging.Blur(&warped, 3)

		res, err := segmentation.ShapeWalkMultiple(vision.NewImage(i2), starts, false)
		if err != nil {
			return err
		}

		d.GotDebugImage(res, "shapes")
	}

	return nil
}

func TestChessCheatRed1(t *testing.T) {
	d := vision.NewMultipleImageTestDebugger("chess/boardseliot2", "*.png")
	err := d.Process(&ChessImageProcessDebug{FindChessCornersPinkCheat})
	if err != nil {
		t.Fatal(err)
	}
}
