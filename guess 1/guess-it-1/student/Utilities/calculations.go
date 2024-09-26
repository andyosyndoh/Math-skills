package utilities

import "math"

// Average calculates the mean of the inputs.
func Average(numbers []float64) float64 {
	var total float64
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	output := total / float64(len(numbers))

	return (output)
}

// Standarddev calculates the standard deviation of the inputs.
func StandardDev(numbers []float64) float64 {
	new := Variance(numbers)
	result := math.Sqrt(new)
	return (result)
}

//Variance calculates the variance of the inputs.
func Variance(numbers []float64) float64 {
	var differences []float64
	for i := 0; i < len(numbers); i++ {
		dif := numbers[i] - Average(numbers)
		differences = append(differences, dif*dif)
	}
	result := Average(differences)
	return (result)
}
