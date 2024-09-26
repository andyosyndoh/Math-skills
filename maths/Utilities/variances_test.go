package Utilities

import (
	"math"
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(2)
	output := math.Round(Variance(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}
