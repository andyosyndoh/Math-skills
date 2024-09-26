package utilities

// LinearRegression calculates the slope and intercept of the best-fit line
func LinearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))

	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n

	return m, b
}
