package Utilities

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(3)
	output := math.Round(Average(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}
