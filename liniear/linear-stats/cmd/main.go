package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"mathskills/utilities"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid Number of Arguments")
	}
	arg := "../" + os.Args[1]

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
	var index []float64
	for i, line := range new {
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println("Error in Parsing number:", err)
			return
		}
		numbers = append(numbers, float64(num))
		index = append(index, float64(i))
	}

	// calling functions in the directory utilities
	m, c := utilities.LinearRegression(index, numbers)
	Pc := utilities.PearsonCorrelation(index, numbers)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", Pc)
}
