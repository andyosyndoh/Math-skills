package main

import (
	"bufio"
	"fmt"
	"math"
	utilities "mathskills/Utilities"
	"os"
	"strconv"
)

func main() {
	last := 40
	numbers := []float64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error in Parsing number:", err)
			return
		}
		numbers = append(numbers, (num))

		if len(numbers) > last {
			numbers = numbers[len(numbers)-last:]
		}

		if len(numbers) > 1 {
			high, low := Solve(numbers)
			fmt.Printf("%v %v\n", low, high)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

// Solve takes the available inputs and calculate the highest and lowest range.
func Solve(numbers []float64) (int, int) {
	mean := utilities.Average(numbers)
	sd := utilities.StandardDev(numbers)
	high := int(math.Round(mean + 3*sd))
	low := int(math.Round(mean - 3*sd))
	return high, low
}
