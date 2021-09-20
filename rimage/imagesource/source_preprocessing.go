package imagesource

import (
	"context"
	"image"

	"github.com/go-errors/errors"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"

	"go.viam.com/core/camera"
	"go.viam.com/core/config"
	"go.viam.com/core/registry"
	"go.viam.com/core/rimage"
	"go.viam.com/core/robot"
)

func init() {
	registry.RegisterCamera("preprocessDepth", registry.Camera{Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (camera.Camera, error) {
		return newPreprocessDepth(r, config)
	}})
}

// PreprocessDepthSource applies pre-processing functions to depth maps in order to smooth edges and fill holes.
type PreprocessDepthSource struct {
	source gostream.ImageSource
}

// Close closes the source
func (os *PreprocessDepthSource) Close() error {
	return nil
}

// Next applies depth preprocessing to the next image
func (os *PreprocessDepthSource) Next(ctx context.Context) (image.Image, func(), error) {
	i, closer, err := os.source.Next(ctx)
	if err != nil {
		return i, closer, err
	}
	defer closer()
	ii := rimage.ConvertToImageWithDepth(i)
	if ii.Depth == nil {
		return nil, nil, errors.New("no depth")
	}
	ii, err = rimage.PreprocessDepthMap(ii)
	if ii.Depth == nil {
		return nil, nil, err
	}
	return ii, func() {}, nil
}

func newPreprocessDepth(r robot.Robot, config config.Component) (camera.Camera, error) {
	source, ok := r.CameraByName(config.Attributes.String("source"))
	if !ok {
		return nil, errors.Errorf("cannot find source camera (%s)", config.Attributes.String("source"))
	}
	return &camera.ImageSource{&PreprocessDepthSource{source}}, nil

}
