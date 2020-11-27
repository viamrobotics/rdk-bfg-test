package ml

import (
	"testing"
)

func _makeSimpleTest() ([][]float64, []int) {
	data := [][]float64{
		[]float64{0, 0, 10},
		[]float64{1, 0, 11},
		[]float64{0, 1, 12},
		[]float64{1, 1, 13},
		[]float64{0, 0, 14},
		[]float64{1, 0, 15},
		[]float64{0, 1, 16},
		[]float64{1, 1, 17},
	}

	correct := []int{
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
	}

	return data, correct
}

func _checkCorrectness(t *testing.T, desc string, theClassifier Classifier, data [][]float64, correct []int) {
	for x, row := range data {
		got, err := theClassifier.Classify(row)
		if err != nil {
			t.Fatal(err)
		}
		if got != correct[x] {
			t.Errorf("wrong result for row %d, data: %v correct: %v got: %v", x, row, correct[x], got)
		}
	}

}

func TestGLSimple1(t *testing.T) {
	data, correct := _makeSimpleTest()

	c := &GoLearnClassifier{}
	err := c.Train(data, correct)
	if err != nil {
		panic(err)
	}

	_checkCorrectness(t, "TestGLSimple1", c, data, correct)
}

func TestGLSimpleNN1(t *testing.T) {
	data, correct := _makeSimpleTest()

	c := &GoLearnNNClassifier{}
	err := c.Train(data, correct)
	if err != nil {
		panic(err)
	}

	_checkCorrectness(t, "TestGLSimpleNN1 - a", c, data, correct)

	for _, r := range data {
		r[2] = r[2] + 1
	}

	_checkCorrectness(t, "TestGLSimpleNN1 - b", c, data, correct)
}
