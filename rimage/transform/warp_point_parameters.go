package transform

import (
	"image"
	"image/color"
	"math"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	"github.com/pkg/errors"

	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/rimage"
)

// DepthColorWarpTransforms TODO.
type DepthColorWarpTransforms struct {
	ColorTransform, DepthTransform rimage.TransformationMatrix
	*AlignConfig                   // anonymous fields
}

// ImagePointTo3DPoint takes in a image coordinate and returns the 3D point from the warp points.
func (dct *DepthColorWarpTransforms) ImagePointTo3DPoint(point image.Point, d rimage.Depth) (r3.Vector, error) {
	i, j := float64(point.X-dct.OutputOrigin.X), float64(point.Y-dct.OutputOrigin.Y)
	return r3.Vector{i, j, float64(d)}, nil
}

// ImageWithDepthToPointCloud TODO.
func (dct *DepthColorWarpTransforms) ImageWithDepthToPointCloud(
	ii *rimage.ImageWithDepth,
	crop ...image.Rectangle) (pointcloud.PointCloud, error) {
	if ii.Depth == nil {
		return nil, errors.New("image with depth has no depth channel. Cannot project to Pointcloud")
	}
	var rect *image.Rectangle
	if len(crop) > 1 {
		return nil, errors.Errorf("cannot have more than one cropping rectangle, got %v", crop)
	}
	if len(crop) == 1 {
		rect = &crop[0]
	}
	var iwd *rimage.ImageWithDepth
	var err error
	if ii.IsAligned() {
		iwd = ii
	} else {
		iwd, err = dct.AlignColorAndDepthImage(ii.Color, ii.Depth)
		if err != nil {
			return nil, err
		}
	}
	// Check dimensions, they should be equal between the color and depth frame
	if iwd.Depth.Width() != iwd.Color.Width() || iwd.Depth.Height() != iwd.Color.Height() {
		return nil, errors.Errorf("depth map and color dimensions don't match Depth(%d,%d) != Color(%d,%d)",
			iwd.Depth.Width(), iwd.Depth.Height(), iwd.Color.Width(), iwd.Color.Height())
	}
	// All points now in Common frame
	pc := pointcloud.New()

	startX, startY := 0, 0
	endX, endY := iwd.Width(), iwd.Height()
	// if optional crop rectangle is provided, use intersections of rectangle and image window and iterate through it
	if rect != nil {
		newBounds := rect.Intersect(iwd.Bounds())
		startX, startY = newBounds.Min.X, newBounds.Min.Y
		endX, endY = newBounds.Max.X, newBounds.Max.Y
	}
	// TODO (bijan): this is a naive projection to 3D space, implement a better algo for warp points
	// Will need more than 2 points for warp points to create better projection
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			z := iwd.Depth.GetDepth(x, y)
			if z == 0 {
				continue
			}
			c := iwd.Color.GetXY(x, y)
			r, g, b := c.RGB255()
			i, j := float64(x-dct.OutputOrigin.X), float64(y-dct.OutputOrigin.Y)
			err := pc.Set(pointcloud.NewVector(i, j, float64(z)), pointcloud.NewColoredData(color.NRGBA{r, g, b, 255}))
			if err != nil {
				return nil, err
			}
		}
	}
	return pc, nil
}

// AlignColorAndDepthImage will warp the color and depth map in order to have them aligned on top of each other.
func (dct *DepthColorWarpTransforms) AlignColorAndDepthImage(col *rimage.Image, dep *rimage.DepthMap) (*rimage.ImageWithDepth, error) {
	if col == nil {
		return nil, errors.New("no color image present to align")
	}
	if dep == nil {
		return nil, errors.New("no depth image present to align")
	}
	if col.Width() != dct.ColorInputSize.X ||
		col.Height() != dct.ColorInputSize.Y ||
		dep.Width() != dct.DepthInputSize.X ||
		dep.Height() != dct.DepthInputSize.Y {
		return nil, errors.Errorf("unexpected aligned dimensions c:(%d,%d) d:(%d,%d) config: %#v",
			col.Width(), col.Height(), dep.Width(), dep.Height(), dct.AlignConfig)
	}

	c2 := rimage.WarpImage(col, dct.ColorTransform, dct.OutputSize)
	dm2 := dep.Warp(dct.DepthTransform, dct.OutputSize)

	return rimage.MakeImageWithDepth(c2, dm2, true), nil
}

// PointCloudToImageWithDepth takes a PointCloud with color info and returns an ImageWithDepth from the
// perspective of the color camera referenceframe.
func (dct *DepthColorWarpTransforms) PointCloudToImageWithDepth(
	cloud pointcloud.PointCloud,
) (*rimage.ImageWithDepth, error) {
	meta := cloud.MetaData()
	// Needs to be a pointcloud with color
	if !meta.HasColor {
		return nil, errors.New("pointcloud has no color information, cannot create an image with depth")
	}
	// ImageWithDepth will be in the camera frame of the RGB camera.
	// Points outside of the frame will be discarded.
	// Assumption is that points in pointcloud are in mm.
	width, height := dct.OutputSize.X, dct.OutputSize.Y
	color := rimage.NewImage(width, height)
	depth := rimage.NewEmptyDepthMap(width, height)
	// TODO(bijan): naive implementation until we get get more points in the warp config
	cloud.Iterate(0, 0, func(pt r3.Vector, d pointcloud.Data) bool {
		j := pt.X - meta.MinX
		i := pt.Y - meta.MinY
		x, y := int(math.Round(j)), int(math.Round(i))
		z := int(pt.Z)
		// if point has color and is inside the RGB image bounds, add it to the images
		if x >= 0 && x < width && y >= 0 && y < height && d != nil && d.HasColor() {
			r, g, b := d.RGB255()
			color.Set(image.Point{x, y}, rimage.NewColor(r, g, b))
			depth.Set(x, y, rimage.Depth(z))
		}
		return true
	})
	return rimage.MakeImageWithDepth(color, depth, true), nil
}

// NewDepthColorWarpTransforms TODO.
func NewDepthColorWarpTransforms(config *AlignConfig, logger golog.Logger) (*DepthColorWarpTransforms, error) {
	var err error
	dst := rimage.ArrayToPoints([]image.Point{{0, 0}, {config.OutputSize.X, config.OutputSize.Y}})

	if config.WarpFromCommon {
		config, err = config.ComputeWarpFromCommon(logger)
		if err != nil {
			return nil, err
		}
	}

	colorPoints := rimage.ArrayToPoints(config.ColorWarpPoints)
	depthPoints := rimage.ArrayToPoints(config.DepthWarpPoints)

	colorTransform := rimage.GetPerspectiveTransform(colorPoints, dst)
	depthTransform := rimage.GetPerspectiveTransform(depthPoints, dst)

	return &DepthColorWarpTransforms{colorTransform, depthTransform, config}, nil
}
