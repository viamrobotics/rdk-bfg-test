package chess

import (
	"image"
	"testing"

	"go.viam.com/robotcore/rimage"
	"go.viam.com/robotcore/vision/segmentation"
)

type P func(*rimage.ImageWithDepth) (image.Image, []image.Point, error)

type ChessImageProcessDebug struct {
	p P
}

func (dd ChessImageProcessDebug) Process(d *rimage.MultipleImageTestDebugger, fn string, img image.Image) error {
	out, corners, err := dd.p(rimage.ConvertToImageWithDepth(img))
	if err != nil {
		return err
	}

	swOptions := segmentation.ShapeWalkOptions{}
	swOptions.MaxRadius = 50

	d.GotDebugImage(out, "corners")

	if corners != nil {
		warped, err := warpColorAndDepthToChess(rimage.ConvertToImageWithDepth(img), corners)
		if err != nil {
			return err
		}

		d.GotDebugImage(warped.Color, "warped")

		starts := []image.Point{}
		for x := 50; x <= 750; x += 100 {
			for y := 50; y <= 750; y += 100 {
				starts = append(starts, image.Point{x, y})
			}
		}

		res, err := segmentation.ShapeWalkMultiple(warped.Color, starts, swOptions)
		if err != nil {
			return err
		}

		d.GotDebugImage(res, "shapes")

		if true {
			out := rimage.NewImageFromBounds(res.Bounds())
			for idx, p := range starts {
				count := res.PixelsInSegmemnt(idx + 1)
				clr := rimage.Red

				if count > 7000 {
					clr = rimage.Green
				}

				out.Circle(p, 20, clr)

			}

			d.GotDebugImage(out, "marked")
		}

		if false {
			clusters, err := rimage.ClusterFromImage(warped.Color, 4)
			if err != nil {
				return err
			}

			clustered := rimage.ClusterImage(clusters, warped.Color)

			d.GotDebugImage(clustered, "kmeans")
		}
	}

	return nil
}

func TestChessCheatRed1(t *testing.T) {
	d := rimage.NewMultipleImageTestDebugger(t, "chess/boardseliot2", "*.png")
	err := d.Process(&ChessImageProcessDebug{FindChessCornersPinkCheat})
	if err != nil {
		t.Fatal(err)
	}
}
