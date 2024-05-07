package Utilities

// function for calculating variances
func Variance(numbers []float64) float64 {
	var differences []float64
	for i := 0; i< len(numbers); i++ {
		dif := numbers[i] - Average(numbers)
		differences = append(differences, dif*dif)
	}
	result := Average(differences)
	return (result)
}