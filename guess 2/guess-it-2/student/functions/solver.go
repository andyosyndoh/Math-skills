package utilities

import "math"

func Solve(Indexes, values []float64, xmean, ymean float64) (int, int) {
	m, b := LinearRegression(Indexes, values, xmean, ymean)

	y := m*(float64(len(Indexes))) + b

	low := y - (3 * StandardDev(values))
	high := y + (3 * StandardDev(values))

	return int(math.Round(low)), int(math.Round(high))
}
