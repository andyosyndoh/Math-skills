package Utilities

import (
	"sort"
)

// function for calculating the median
func Median(numbers []float64) float64 {
	sort.Float64s(numbers)
	var mid float64
	if len(numbers)%2 == 0 {
		first := len(numbers) / 2
		second := (len(numbers) / 2) - 1
		mid = (numbers[first] + numbers[second]) / 2
	} else if len(numbers)%2 == 1 {
		first := (len(numbers) - 1) / 2
		mid = numbers[first]
	}
	return (mid)
}
