package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mathskills "mathskills/Utilities"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid Number of Arguments")
	}
	arg := os.Args[1]

	// if the data.txt file is missing
	_, err := os.Open(arg)
	if err != nil {
		fmt.Println("error")
		return
	}

	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	new := strings.Split(string(file), "\n")
	var numbers []float64
	for _, line := range new {
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Error in Parsing number:", err)
			return
		}
		numbers = append(numbers, float64(num))
	}

	// calling functions in the directory utilities
	output := mathskills.Average(numbers)
	fmt.Println("Average:", int(math.Round(output)))
	output2 := mathskills.Median(numbers)
	fmt.Println("Median:", int(math.Round(output2)))
	output3 := mathskills.Variance(numbers)
	fmt.Println("Variance:", int(math.Round(output3)))
	output4 := mathskills.StandardDev(numbers)
	fmt.Println("Standard Deviation:", int(math.Round(output4)))
}
