package robot

import (
	"testing"
)

func TestConfigRobot(t *testing.T) {
	cfg, err := ReadConfig("data/robot.json")
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Components) != 4 {
		t.Errorf("bad config read %v", cfg)
	}

}

func TestConfig1(t *testing.T) {
	cfg, err := ReadConfig("data/cfgtest1.json")
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewRobot(cfg)
	if err != nil {
		t.Fatal(err)
	}

	pic, _, err := r.Cameras[0].NextMatDepthPair()
	if err != nil {
		t.Fatal(err)
	}
	defer pic.Close()

	if pic.Cols() < 100 {
		t.Errorf("pictures seems wrong %d %d", pic.Cols(), pic.Rows())
	}

}

func TestConfigDummy(t *testing.T) {
	cfg, err := ReadConfig("data/dummy.json")
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewRobot(cfg)
	if err != nil {
		t.Fatal(err)
	}
	r.Close()
}
