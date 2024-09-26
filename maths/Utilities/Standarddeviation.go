package Utilities

import "math"

// function for calculating the standard deviations
func StandardDev(numbers []float64) float64 {
	new := Variance(numbers)
	result := math.Sqrt(new)
	return (result)
}
