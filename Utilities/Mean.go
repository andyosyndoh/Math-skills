package Utilities

// function for calculating the average
func Average(numbers []float64) float64 {
	var total float64
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	output := total / float64(len(numbers))

	return (output)
}
