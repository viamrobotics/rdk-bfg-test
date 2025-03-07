package tflitecpu

import (
	"context"
	"math/rand"
	"net"
	"reflect"
	"testing"

	"github.com/edaniels/golog"
	"github.com/nfnt/resize"
	"go.viam.com/test"
	"go.viam.com/utils/artifact"
	"go.viam.com/utils/rpc"

	viamgrpc "go.viam.com/rdk/grpc"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage"
	"go.viam.com/rdk/services/mlmodel"
	"go.viam.com/rdk/subtype"
)

func TestEmptyTFLiteConfig(t *testing.T) {
	ctx := context.Background()
	emptyCfg := TFLiteConfig{} // empty config

	// Test that empty config gives error about loading model
	emptyGot, err := NewTFLiteCPUModel(ctx, &emptyCfg, "fakeModel")
	test.That(t, emptyGot, test.ShouldBeNil)
	test.That(t, err.Error(), test.ShouldContainSubstring, "could not add model")
}

func TestTFLiteCPUDetector(t *testing.T) {
	ctx := context.Background()
	modelLoc := artifact.MustPath("vision/tflite/effdet0.tflite")
	cfg := TFLiteConfig{ // detector config
		ModelPath:  modelLoc,
		NumThreads: 2,
	}
	// Test that a detector would give the expected output on the dog image
	// Creating the model should populate model and attrs, but not metadata
	out, err := NewTFLiteCPUModel(ctx, &cfg, "myDet")
	got := out.(*Model)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, got.model, test.ShouldNotBeNil)
	test.That(t, got.attrs, test.ShouldNotBeNil)
	test.That(t, got.metadata, test.ShouldBeNil)

	// Test that the Metadata() works on detector
	gotMD, err := got.Metadata(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotMD, test.ShouldNotBeNil)
	test.That(t, got.metadata, test.ShouldNotBeNil)

	test.That(t, gotMD.Inputs[0].Name, test.ShouldResemble, "image")
	test.That(t, gotMD.Outputs[0].Name, test.ShouldResemble, "location")
	test.That(t, gotMD.Outputs[1].Name, test.ShouldResemble, "category")
	test.That(t, gotMD.Outputs[2].Name, test.ShouldResemble, "score")
	test.That(t, gotMD.Inputs[0].DataType, test.ShouldResemble, "uint8")
	test.That(t, gotMD.Outputs[0].DataType, test.ShouldResemble, "float32")
	test.That(t, gotMD.Outputs[1].AssociatedFiles[0].Name, test.ShouldResemble, "labelmap.txt")

	// Test that the Infer() works on detector
	pic, err := rimage.NewImageFromFile(artifact.MustPath("vision/tflite/dogscute.jpeg"))
	test.That(t, err, test.ShouldBeNil)
	resized := resize.Resize(uint(got.metadata.Inputs[0].Shape[1]), uint(got.metadata.Inputs[0].Shape[2]), pic, resize.Bilinear)
	imgBytes := rimage.ImageToUInt8Buffer(resized)
	test.That(t, imgBytes, test.ShouldNotBeNil)
	inputMap := make(map[string]interface{})
	inputMap["image"] = imgBytes

	gotOutput, err := got.Infer(ctx, inputMap)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotOutput, test.ShouldNotBeNil)

	test.That(t, gotOutput["number of detections"], test.ShouldResemble, []float32{25})
	test.That(t, len(gotOutput["score"].([]float32)), test.ShouldResemble, 25)
	test.That(t, len(gotOutput["location"].([]float32)), test.ShouldResemble, 100)
	test.That(t, len(gotOutput["category"].([]float32)), test.ShouldResemble, 25)
	test.That(t, gotOutput["category"].([]float32)[0], test.ShouldEqual, 17) // 17 is dog
}

func TestTFLiteCPUClassifier(t *testing.T) {
	ctx := context.Background()
	modelLoc := artifact.MustPath("vision/tflite/effnet0.tflite")
	cfg := TFLiteConfig{ // classifier config
		ModelPath:  modelLoc,
		NumThreads: 2,
	}

	// Test that the tflite classifier gives the expected output on the lion image
	out, err := NewTFLiteCPUModel(ctx, &cfg, "myClass")
	got := out.(*Model)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, got.model, test.ShouldNotBeNil)
	test.That(t, got.attrs, test.ShouldNotBeNil)
	test.That(t, got.metadata, test.ShouldBeNil)

	// Test that the Metadata() works on classifier
	gotMD, err := got.Metadata(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotMD, test.ShouldNotBeNil)

	test.That(t, gotMD.Inputs[0].Name, test.ShouldResemble, "image")
	test.That(t, gotMD.Outputs[0].Name, test.ShouldResemble, "probability")
	test.That(t, gotMD.Inputs[0].DataType, test.ShouldResemble, "uint8")
	test.That(t, gotMD.Outputs[0].DataType, test.ShouldResemble, "uint8")
	test.That(t, gotMD.Outputs[0].AssociatedFiles[0].Name, test.ShouldContainSubstring, ".txt")
	test.That(t, gotMD.Outputs[0].AssociatedFiles[0].LabelType, test.ShouldResemble, mlmodel.LabelTypeTensorAxis)

	// Test that the Infer() works on a classifier
	pic, err := rimage.NewImageFromFile(artifact.MustPath("vision/tflite/lion.jpeg"))
	test.That(t, err, test.ShouldBeNil)
	resized := resize.Resize(uint(got.metadata.Inputs[0].Shape[1]), uint(got.metadata.Inputs[0].Shape[2]), pic, resize.Bilinear)
	imgBytes := rimage.ImageToUInt8Buffer(resized)
	test.That(t, imgBytes, test.ShouldNotBeNil)
	inputMap := make(map[string]interface{})
	inputMap["image"] = imgBytes

	gotOutput, err := got.Infer(ctx, inputMap)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotOutput, test.ShouldNotBeNil)

	test.That(t, gotOutput["probability"].([]uint8), test.ShouldNotBeNil)
	test.That(t, gotOutput["probability"].([]uint8)[290], test.ShouldEqual, 0)
	test.That(t, gotOutput["probability"].([]uint8)[291], test.ShouldBeGreaterThan, 200) // 291 is lion
	test.That(t, gotOutput["probability"].([]uint8)[292], test.ShouldEqual, 0)
}

func TestTFLiteCPUTextModel(t *testing.T) {
	// Setup
	ctx := context.Background()
	modelLoc := artifact.MustPath("vision/tflite/mobileBERT.tflite")

	cfg := TFLiteConfig{ // text classifier config
		ModelPath:  modelLoc,
		NumThreads: 1,
	}

	// Test that even a text classifier gives an output with good input
	out, err := NewTFLiteCPUModel(ctx, &cfg, "myTextModel")
	got := out.(*Model)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, got.model, test.ShouldNotBeNil)
	test.That(t, got.attrs, test.ShouldNotBeNil)
	test.That(t, got.metadata, test.ShouldBeNil)

	// Test that the Metadata() errors well when metadata does not exist
	_, err = got.Metadata(ctx)
	test.That(t, err.Error(), test.ShouldContainSubstring, "metadata does not exist")
	test.That(t, got.metadata, test.ShouldBeNil)

	// Test that the Infer() works even on a text classifier
	inputMap := make(map[string]interface{})
	inputMap["text"] = makeRandomSlice(got.model.Info.InputHeight)
	test.That(t, len(inputMap["text"].([]int32)), test.ShouldEqual, 384)
	gotOutput, err := got.Infer(ctx, inputMap)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotOutput, test.ShouldNotBeNil)

	test.That(t, len(gotOutput), test.ShouldEqual, 2)
	test.That(t, gotOutput["output0"], test.ShouldNotBeNil)
	test.That(t, gotOutput["output1"], test.ShouldNotBeNil)
	test.That(t, gotOutput["output2"], test.ShouldBeNil)

	test.That(t, len(gotOutput["output0"].([]float32)), test.ShouldEqual, 384)
	test.That(t, len(gotOutput["output1"].([]float32)), test.ShouldEqual, 384)
}

func TestTFLiteCPUClient(t *testing.T) {
	logger := golog.NewTestLogger(t)
	listener1, err := net.Listen("tcp", "localhost:0")
	test.That(t, err, test.ShouldBeNil)
	rpcServer, err := rpc.NewServer(logger, rpc.WithUnauthenticated())
	test.That(t, err, test.ShouldBeNil)

	modelParams := TFLiteConfig{ // classifier config
		ModelPath:  artifact.MustPath("vision/tflite/effdet0.tflite"),
		NumThreads: 2,
	}
	myModel, err := NewTFLiteCPUModel(context.Background(), &modelParams, "myModel")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, myModel, test.ShouldNotBeNil)

	omMap := map[resource.Name]interface{}{
		mlmodel.Named("testName"): myModel,
	}
	svc, err := subtype.New(omMap)
	test.That(t, err, test.ShouldBeNil)
	resourceSubtype := registry.ResourceSubtypeLookup(mlmodel.Subtype)
	resourceSubtype.RegisterRPCService(context.Background(), rpcServer, svc)

	go rpcServer.Serve(listener1)
	defer rpcServer.Stop()

	// Prep img
	pic, err := rimage.NewImageFromFile(artifact.MustPath("vision/tflite/dogscute.jpeg"))
	test.That(t, err, test.ShouldBeNil)
	resized := resize.Resize(320, 320, pic, resize.Bilinear)
	imgBytes := rimage.ImageToUInt8Buffer(resized)
	test.That(t, imgBytes, test.ShouldNotBeNil)
	inputMap := make(map[string]interface{})
	inputMap["image"] = imgBytes

	conn, err := viamgrpc.Dial(context.Background(), listener1.Addr().String(), logger)
	test.That(t, err, test.ShouldBeNil)
	client := mlmodel.NewClientFromConn(context.Background(), conn, "testName", logger)
	// Test call to Metadata
	gotMD, err := client.Metadata(context.Background())
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotMD, test.ShouldNotBeNil)
	test.That(t, gotMD.ModelType, test.ShouldEqual, "tflite_detector")
	test.That(t, gotMD.Inputs[0].Name, test.ShouldResemble, "image")
	test.That(t, gotMD.Outputs[0].Name, test.ShouldResemble, "location")
	test.That(t, gotMD.Outputs[1].Name, test.ShouldResemble, "category")
	test.That(t, gotMD.Outputs[2].Name, test.ShouldResemble, "score")
	test.That(t, gotMD.Outputs[3].Name, test.ShouldResemble, "number of detections")
	test.That(t, gotMD.Outputs[1].Description, test.ShouldContainSubstring, "categories of the detected boxes")
	test.That(t, gotMD.Inputs[0].DataType, test.ShouldResemble, "uint8")
	test.That(t, gotMD.Outputs[0].DataType, test.ShouldResemble, "float32")
	test.That(t, gotMD.Outputs[1].AssociatedFiles[0].Name, test.ShouldResemble, "labelmap.txt")

	// Test call to Infer
	gotOutput, err := client.Infer(context.Background(), inputMap)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, gotOutput, test.ShouldNotBeNil)
	test.That(t, len(gotOutput), test.ShouldEqual, 4)
	locs := reflect.ValueOf(gotOutput["location"])
	test.That(t, locs.Len(), test.ShouldEqual, 100)
	scores := reflect.ValueOf(gotOutput["score"])
	test.That(t, scores.Len(), test.ShouldEqual, 25)
	nDets := reflect.ValueOf(gotOutput["number of detections"])
	test.That(t, nDets.Len(), test.ShouldEqual, 1)
	test.That(t, nDets.Index(0).Interface().(float64), test.ShouldResemble, float64(25))
	test.That(t, reflect.TypeOf(gotOutput["category"]).Kind(), test.ShouldResemble, reflect.Slice)
	cats := reflect.ValueOf(gotOutput["category"])
	test.That(t, cats.Len(), test.ShouldEqual, 25)
	test.That(t, cats.Index(0).Interface().(float64), test.ShouldResemble, float64(17)) // 17 is dog
}

func makeRandomSlice(length int) []int32 {
	out := make([]int32, 0, length)
	for i := 0; i < length; i++ {
		x := rand.Int31n(100)
		out = append(out, x)
	}
	return out
}
