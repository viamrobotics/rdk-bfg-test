package vision

import (
	"fmt"
	"image"
	"sort"

	"gocv.io/x/gocv"
)

func WarpColorAndDepthToChess(color, depth gocv.Mat, corners []image.Point) (gocv.Mat, gocv.Mat, error) {
	dst := []image.Point{
		image.Pt(0, 800),
		image.Pt(0, 0),
		image.Pt(800, 800),
		image.Pt(800, 0),
	}

	m := gocv.GetPerspectiveTransform(corners, dst)
	defer m.Close()

	warped := gocv.NewMat()
	gocv.WarpPerspective(color, &warped, m, image.Point{800, 800})

	warpedDepth := gocv.Mat{}
	if !depth.Empty() {
		warpedDepth = gocv.NewMatWithSize(800, 800, depth.Type())
		gocv.WarpPerspective(depth, &warpedDepth, m, image.Point{800, 800})
	}

	return warped, warpedDepth, nil
}

// returns point in a1, a8, h1, h8 order
func FindChessCorners(img gocv.Mat) ([]image.Point, error) {
	a1Corner := image.Point{145, 45}
	a8Corner := image.Point{520, 52}
	h1Corner := image.Point{125, 440}
	h8Corner := image.Point{545, 440}

	return []image.Point{a1Corner, a8Corner, h1Corner, h8Corner}, nil
}

func getMinChessCorner(chess string) image.Point {
	var x = int(chess[0]-'A') * 100
	var y = 100 * (7 - int(chess[1]-'1'))
	return image.Point{x, y}
}

func GetChessPieceHeight(square string, warpedDepth gocv.Mat) float64 {
	if warpedDepth.Cols() != 800 || warpedDepth.Rows() != 800 {
		panic("bad image size pased to GetChessPieceHeight")
	}
	data := []float64{}

	corner := getMinChessCorner(square)
	for x := corner.X + 33; x < corner.X+66; x++ {
		for y := corner.Y + 33; y < corner.Y+66; y++ {
			d := warpedDepth.GetDoubleAt(y, x)
			data = append(data, d)
		}
	}

	sort.Float64s(data)

	for idx, f := range data {
		fmt.Printf("%d %f\n", idx, f)
	}
	fmt.Println("--------")

	// since there is some noise, let's try and remove the outliers
	// TODO: make this smarter
	takeOff := len(data) / 10
	data = data[takeOff:]
	data = data[0 : len(data)-takeOff]

	min := data[0]
	max := data[len(data)-1]

	return max - min
}
