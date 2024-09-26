package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	utilities "guess-it/functions"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	values := []float64{}
	Indexes := []float64{}

	index := 0

	for scanner.Scan() {
		input := scanner.Text()

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("error during conversion")
			continue
		}
		values = append(values, num)
		Indexes = append(Indexes, float64(index))

		xmean := utilities.Average(Indexes)
		ymean := utilities.Average(values)
		if len(values) > 1 {
			m, b := utilities.Solve(Indexes, values, xmean, ymean)
			fmt.Printf("%v %v\n", m, b)
		}
		index++
	}
}
