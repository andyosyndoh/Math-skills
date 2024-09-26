package utilities

import "math"

// Average calculates the mean of the inputs.
func Average(numbers []float64) float64 {
	var total float64
	for _, ch := range numbers {
		total += ch
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

// Variance calculates the variance of the inputs.
func Variance(numbers []float64) float64 {
	var differences float64
	avg := Average(numbers)
	for _, ch := range numbers {
		differences += (ch - avg) * (ch - avg)
	}
	// result := Average(differences)
	return differences / float64(len(numbers))
}

// LinearRegression calculates the slope and intercept of the best-fit line
func LinearRegression(x, y []float64, xmean, ymean float64) (float64, float64) {
	// n := float64(len(x))

	var sumXY, sumX2 float64

	for i := 0; i < len(x); i++ {

		sumXY += (x[i] - xmean) * (y[i] - ymean)
		sumX2 += (x[i] - xmean) * (x[i] - xmean)
	}

	m := sumXY / sumX2
	b := ymean - (m * xmean)

	return m, b
}
