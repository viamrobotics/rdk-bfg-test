// Package ffmpeg provides an implementation for an ffmpeg based camera
package ffmpeg

import (
	"context"
	"image"
	"image/jpeg"
	"io"
	"os/exec"
	"sync"
	"sync/atomic"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapio"
	viamutils "go.viam.com/utils"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage/transform"
	"go.viam.com/rdk/utils"
)

// AttrConfig is the attribute struct for ffmpeg cameras.
type AttrConfig struct {
	CameraParameters     *transform.PinholeCameraIntrinsics `json:"intrinsic_parameters,omitempty"`
	DistortionParameters *transform.BrownConrady            `json:"distortion_parameters,omitempty"`
	Debug                bool                               `json:"debug,omitempty"`
	VideoPath            string                             `json:"video_path"`
	InputKWArgs          map[string]interface{}             `json:"input_kw_args,omitempty"`
	Filters              []FilterAttrs                      `json:"filters,omitempty"`
	OutputKWArgs         map[string]interface{}             `json:"output_kw_args,omitempty"`
}

// FilterAttrs is a struct to used to configure ffmpeg filters.
type FilterAttrs struct {
	Name   string                 `json:"name"`
	Args   []string               `json:"args"`
	KWArgs map[string]interface{} `json:"kw_args"`
}

var model = resource.NewDefaultModel("ffmpeg")

func init() {
	registry.RegisterComponent(camera.Subtype, model, registry.Component{
		Constructor: func(ctx context.Context, _ registry.Dependencies, cfg config.Component, logger golog.Logger) (interface{}, error) {
			attrs, ok := cfg.ConvertedAttributes.(*AttrConfig)
			if !ok {
				return nil, utils.NewUnexpectedTypeError(attrs, cfg.ConvertedAttributes)
			}
			return NewFFMPEGCamera(ctx, attrs, logger)
		},
	})

	config.RegisterComponentAttributeMapConverter(
		camera.Subtype,
		model,
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf AttrConfig
			attrs, err := config.TransformAttributeMapToStruct(&conf, attributes)
			if err != nil {
				return nil, err
			}
			result, ok := attrs.(*AttrConfig)
			if !ok {
				return nil, utils.NewUnexpectedTypeError(result, attrs)
			}
			return result, nil
		},
		&AttrConfig{},
	)
}

type ffmpegCamera struct {
	gostream.VideoReader
	cancelFunc              context.CancelFunc
	activeBackgroundWorkers sync.WaitGroup
	inClose                 func() error
	outClose                func() error
}

// NewFFMPEGCamera instantiates a new camera which leverages ffmpeg to handle a variety of potential video types.
func NewFFMPEGCamera(ctx context.Context, attrs *AttrConfig, logger golog.Logger) (camera.Camera, error) {
	// make sure ffmpeg is in the path before doing anything else
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return nil, err
	}
	// parse attributes into ffmpeg keyword maps
	outArgs := make(map[string]interface{}, len(attrs.OutputKWArgs))
	for key, value := range attrs.OutputKWArgs {
		outArgs[key] = value
	}
	outArgs["update"] = 1        // always interpret the filename as just a filename, not a pattern
	outArgs["format"] = "image2" // select image file muxer, used to write video frames to image files

	// instantiate camera with cancellable context that will be applied to all spawned processes
	cancelableCtx, cancel := context.WithCancel(context.Background())
	ffCam := &ffmpegCamera{cancelFunc: cancel}

	// launch thread to run ffmpeg and pull images from the url and put them into the pipe
	in, out := io.Pipe()

	// Note(erd): For some reason, when running with the race detector, we need to close the pipe
	// even if we kill the process in order for the underlying command Wait to complete.
	ffCam.inClose = in.Close
	ffCam.outClose = out.Close

	writer := &zapio.Writer{Log: logger.Desugar(), Level: zap.DebugLevel}

	ffCam.activeBackgroundWorkers.Add(1)
	viamutils.ManagedGo(func() {
		stream := ffmpeg.Input(attrs.VideoPath, attrs.InputKWArgs)
		for _, filter := range attrs.Filters {
			stream = stream.Filter(filter.Name, filter.Args, filter.KWArgs)
		}
		stream = stream.Output("pipe:", outArgs)
		stream.Context = cancelableCtx
		cmd := stream.WithOutput(out).WithErrorOutput(writer).Compile()
		if err := cmd.Run(); err != nil {
			if viamutils.FilterOutError(err, context.Canceled) == nil ||
				viamutils.FilterOutError(err, context.DeadlineExceeded) == nil {
				return
			}
			if cmd.ProcessState.ExitCode() != 0 {
				panic(err)
			}
		}
	}, func() {
		viamutils.UncheckedErrorFunc(writer.Close)
		cancel()
		ffCam.activeBackgroundWorkers.Done()
	})

	// launch thread to consume images from the pipe and store the latest in shared memory
	gotFirstFrame := make(chan struct{})
	var latestFrame atomic.Value
	var gotFirstFrameOnce bool
	ffCam.activeBackgroundWorkers.Add(1)
	viamutils.ManagedGo(func() {
		for {
			if cancelableCtx.Err() != nil {
				return
			}
			img, err := jpeg.Decode(in)
			if err != nil {
				continue
			}
			latestFrame.Store(img)
			if !gotFirstFrameOnce {
				close(gotFirstFrame)
				gotFirstFrameOnce = true
			}
		}
	}, ffCam.activeBackgroundWorkers.Done)

	// when next image is requested simply load the image from where it is stored in shared memory
	reader := gostream.VideoReaderFunc(func(ctx context.Context) (image.Image, func(), error) {
		select {
		case <-cancelableCtx.Done():
			return nil, nil, cancelableCtx.Err()
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-gotFirstFrame:
		}
		return latestFrame.Load().(image.Image), func() {}, nil
	})

	ffCam.VideoReader = reader
	return camera.NewFromReader(ctx, ffCam, &transform.PinholeCameraModel{PinholeCameraIntrinsics: attrs.CameraParameters}, camera.ColorStream)
}

func (fc *ffmpegCamera) Close(ctx context.Context) error {
	fc.cancelFunc()
	viamutils.UncheckedError(fc.inClose())
	viamutils.UncheckedError(fc.outClose())
	fc.activeBackgroundWorkers.Wait()
	return nil
}
