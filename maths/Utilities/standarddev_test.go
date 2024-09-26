package Utilities

import (
	"math"
	"testing"
)

func TestStandardDev(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(1)
	output := math.Round(StandardDev(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}
