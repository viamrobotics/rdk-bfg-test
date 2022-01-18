package robotimpl

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/a8m/envsubst"
	"github.com/edaniels/golog"
	"github.com/pkg/errors"
	"go.viam.com/test"
	"go.viam.com/utils"

	"go.viam.com/rdk/component/arm"
	"go.viam.com/rdk/component/base"
	"go.viam.com/rdk/component/board"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/metadata/service"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot"
	"go.viam.com/rdk/services/web"
	rdktestutils "go.viam.com/rdk/testutils"
)

func TestRobotReconfigure(t *testing.T) {
	ConfigFromFile := func(t *testing.T, filePath string) *config.Config {
		t.Helper()
		buf, err := envsubst.ReadFile(filePath)
		test.That(t, err, test.ShouldBeNil)
		conf, err := config.FromReader(context.Background(), filePath, bytes.NewReader(buf))
		test.That(t, err, test.ShouldBeNil)
		return conf
	}
	mockSubtype := resource.NewSubtype(resource.ResourceNamespaceRDK, resource.ResourceTypeComponent, resource.SubtypeName("mock"))
	mockNamed := func(name string) resource.Name {
		return resource.NameFromSubtype(mockSubtype, name)
	}
	modelName1 := utils.RandomAlphaString(5)
	modelName2 := utils.RandomAlphaString(5)
	test.That(t, os.Setenv("TEST_MODEL_NAME_1", modelName1), test.ShouldBeNil)
	test.That(t, os.Setenv("TEST_MODEL_NAME_2", modelName2), test.ShouldBeNil)

	registry.RegisterComponent(mockSubtype, modelName1, registry.Component{
		Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (interface{}, error) {
			return &mockFake{x: 5}, nil
		},
	})

	// these settings to be toggled in test cases specifically
	// testing for a reconfigurability mismatch
	reconfigurableTrue := true
	testReconfiguringMismatch := false
	registry.RegisterComponent(mockSubtype, modelName2, registry.Component{
		Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (interface{}, error) {
			if reconfigurableTrue && testReconfiguringMismatch {
				reconfigurableTrue = false
				return &mockFake{x: 5}, nil
			}
			return &mockFake2{x: 5}, nil
		},
	})

	t.Run("no diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")

		ctx := context.Background()
		svc, err := service.New()
		test.That(t, err, test.ShouldBeNil)
		test.That(t, len(svc.All()), test.ShouldEqual, 1)
		ctx = service.ContextWithService(ctx, svc)

		robot, err := New(ctx, conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()
		test.That(t, len(svc.All()), test.ShouldEqual, 7)
		rCopy := make([]resource.Name, 7)
		copy(rCopy, svc.All())

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		test.That(t, robot.Reconfigure(ctx, conf1), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		_, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ResourceByName(board.Named("board1"))
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)

		test.That(t, rCopy, test.ShouldResemble, svc.All())
	})

	t.Run("empty to additive diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		emptyConf := ConfigFromFile(t, "data/diff_config_empty.json")
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}

		ctx := context.Background()
		svc, err := service.New()
		test.That(t, err, test.ShouldBeNil)
		test.That(t, len(svc.All()), test.ShouldEqual, 1)
		ctx = service.ContextWithService(ctx, svc)

		robot, err := New(ctx, emptyConf, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()
		test.That(t, len(svc.All()), test.ShouldEqual, 2)

		test.That(t, robot.Reconfigure(ctx, emptyConf), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ArmNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BoardNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, robot.ResourceNames(), test.ShouldResemble, serviceNames)
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldBeEmpty)

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		test.That(t, robot.Reconfigure(ctx, conf1), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		_, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)

		test.That(t, svc, test.ShouldResemble, service.ContextService(ctx))
		test.That(t, len(svc.All()), test.ShouldEqual, 7)
	})

	t.Run("additive diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		conf4 := ConfigFromFile(t, "data/diff_config_4.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		test.That(t, robot.Reconfigure(context.Background(), conf1), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		baseNames = []resource.Name{base.Named("base1"), base.Named("base2")}
		test.That(t, robot.Reconfigure(context.Background(), conf4), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1", "base2"))
		test.That(t, utils.NewStringSet(robot.BoardNames()...), test.ShouldResemble, utils.NewStringSet("board1"))
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		_, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)
	})

	t.Run("subtractive diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		emptyConf := ConfigFromFile(t, "data/diff_config_empty.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}

		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		_, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		test.That(t, robot.Reconfigure(context.Background(), emptyConf), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ArmNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BoardNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, robot.ResourceNames(), test.ShouldResemble, serviceNames)
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldBeEmpty)

		_, ok = robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeFalse)
	})

	t.Run("modificative diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		conf2 := ConfigFromFile(t, "data/diff_config_2.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		arm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		base1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		board1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		resource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)

		test.That(t, robot.Reconfigure(context.Background(), conf2), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(t, utils.NewStringSet(robot.BoardNames()...), test.ShouldResemble, utils.NewStringSet("board1"))
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 1)

		newArm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newArm1, test.ShouldEqual, arm1)

		newBase1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBase1, test.ShouldEqual, base1)

		newBoard1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBoard1, test.ShouldEqual, board1)

		_, ok = newBoard1.AnalogReaderByName("analog1")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = newBoard1.AnalogReaderByName("analog2")
		test.That(t, ok, test.ShouldBeFalse)

		newResource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newResource1, test.ShouldEqual, resource1)

		newMock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newMock1, test.ShouldEqual, mock1)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)
	})

	t.Run("mixed diff", func(t *testing.T) {
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		conf3 := ConfigFromFile(t, "data/diff_config_3.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		arm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		base1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, base1, test.ShouldNotBeNil)

		board1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		resource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)

		armNames = []resource.Name{arm.Named("arm1")}
		boardNames = []resource.Name{board.Named("board1"), board.Named("board2")}
		baseNames = []resource.Name{base.Named("base2")}
		mockNames = []resource.Name{mockNamed("mock1")}
		test.That(t, robot.Reconfigure(context.Background(), conf3), test.ShouldBeNil)
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base2"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "3"))

		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 1)

		newArm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newArm1, test.ShouldEqual, arm1)

		newBase1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeFalse)
		test.That(t, newBase1, test.ShouldBeNil)

		newBoard1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBoard1, test.ShouldEqual, board1)

		_, ok = robot.BaseByName("base2")
		test.That(t, ok, test.ShouldBeTrue)

		_, ok = robot.BoardByName("board2")
		test.That(t, ok, test.ShouldBeTrue)

		newResource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newResource1, test.ShouldEqual, resource1)

		newMock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newMock1, test.ShouldEqual, mock1)

		_, ok = robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeFalse)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeFalse)
		_, ok = robot.ProcessManager().ProcessByID("3")
		test.That(t, ok, test.ShouldBeTrue)
	})

	t.Run("reconfiguring unreconfigurable", func(t *testing.T) {
		testReconfiguringMismatch = true
		// processing modify will fail
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		conf3 := ConfigFromFile(t, "data/diff_config_4_bad.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		baseNames := []resource.Name{base.Named("base1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		arm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		base1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		board1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		resource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		reconfigurableTrue = false
		err = robot.Reconfigure(context.Background(), conf3)
		test.That(t, err, test.ShouldNotBeNil)
		reconfigurableTrue = true

		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		newArm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newArm1, test.ShouldEqual, arm1)
		newBase1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBase1, test.ShouldEqual, base1)
		newBoard1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBoard1, test.ShouldEqual, board1)
		newResource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newResource1, test.ShouldEqual, resource1)

		newMock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newMock1, test.ShouldEqual, mock1)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)

		testReconfiguringMismatch = false
	})

	t.Run("rollback", func(t *testing.T) {
		// processing modify will fail
		logger := golog.NewTestLogger(t)
		conf1 := ConfigFromFile(t, "data/diff_config_1.json")
		conf3 := ConfigFromFile(t, "data/diff_config_3_bad.json")
		robot, err := New(context.Background(), conf1, logger)
		test.That(t, err, test.ShouldBeNil)
		defer func() {
			test.That(t, robot.Close(context.Background()), test.ShouldBeNil)
		}()

		armNames := []resource.Name{arm.Named("arm1")}
		boardNames := []resource.Name{board.Named("board1")}
		mockNames := []resource.Name{mockNamed("mock1"), mockNamed("mock2")}
		serviceNames := []resource.Name{resource.NameFromSubtype(web.Subtype, "")}
		baseNames := []resource.Name{base.Named("base1")}
		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		arm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)

		base1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)

		board1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)

		resource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)

		mock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock1.(*mockFake).x, test.ShouldEqual, 5)
		test.That(t, mock1.(*mockFake).reconfCount, test.ShouldEqual, 0)

		mock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mock2.(*mockFake2).x, test.ShouldEqual, 5)
		test.That(t, mock2.(*mockFake2).reconfCount, test.ShouldEqual, 0)

		err = robot.Reconfigure(context.Background(), conf3)
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "error processing draft changes")
		test.That(t, err.Error(), test.ShouldContainSubstring, "whoops")

		test.That(t, utils.NewStringSet(robot.RemoteNames()...), test.ShouldBeEmpty)
		test.That(
			t,
			utils.NewStringSet(robot.ArmNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(armNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.GripperNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.CameraNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.BaseNames()...), test.ShouldResemble, utils.NewStringSet("base1"))
		test.That(
			t,
			utils.NewStringSet(robot.BoardNames()...),
			test.ShouldResemble,
			utils.NewStringSet(rdktestutils.ExtractNames(boardNames...)...),
		)
		test.That(t, utils.NewStringSet(robot.SensorNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.ServoNames()...), test.ShouldBeEmpty)
		test.That(t, utils.NewStringSet(robot.FunctionNames()...), test.ShouldBeEmpty)
		test.That(t, rdktestutils.NewResourceNameSet(robot.ResourceNames()...), test.ShouldResemble, rdktestutils.NewResourceNameSet(
			rdktestutils.ConcatResourceNames(
				armNames,
				boardNames,
				mockNames,
				serviceNames,
				baseNames,
			)...))
		test.That(t, utils.NewStringSet(robot.ProcessManager().ProcessIDs()...), test.ShouldResemble, utils.NewStringSet("1", "2"))

		newArm1, ok := robot.ArmByName("arm1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newArm1, test.ShouldEqual, arm1)
		newBase1, ok := robot.BaseByName("base1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBase1, test.ShouldEqual, base1)
		newBoard1, ok := robot.BoardByName("board1")
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newBoard1, test.ShouldEqual, board1)
		newResource1, ok := robot.ResourceByName(arm.Named("arm1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newResource1, test.ShouldEqual, resource1)
		newMock1, ok := robot.ResourceByName(mockNamed("mock1"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newMock1, test.ShouldEqual, mock1)

		newMock2, ok := robot.ResourceByName(mockNamed("mock2"))
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, newMock2, test.ShouldEqual, mock2)

		_, ok = robot.ProcessManager().ProcessByID("1")
		test.That(t, ok, test.ShouldBeTrue)
		_, ok = robot.ProcessManager().ProcessByID("2")
		test.That(t, ok, test.ShouldBeTrue)
	})
}

type mockFake struct {
	x           int
	reconfCount int
}

func (m *mockFake) Reconfigure(ctx context.Context, newResource resource.Reconfigurable) error {
	res, ok := newResource.(*mockFake)
	if !ok {
		return errors.Errorf("expected new arm to be %T but got %T", m, newResource)
	}
	m.x = res.x
	m.reconfCount++
	return nil
}

type mockFake2 struct {
	x           int
	reconfCount int
}
