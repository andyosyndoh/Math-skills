package main

import (
	"math"
	utilities "mathskills/Utilities"
	"testing"
)

func TestMean(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(3)
	output := math.Round(utilities.Average(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}

func TestSolve(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "test", args: args{numbers: []float64{89, 89, 89, 89}}, want: 89, want1: 89},
		{name: "test2", args: args{numbers: []float64{189, 113, 121, 114}}, want: 230, want1: 39},
		{name: "test3", args: args{numbers: []float64{189, 113, 121}}, want: 243, want1: 39},
		{name: "test4", args: args{numbers: []float64{189}}, want: 189, want1: 189},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Solve(tt.args.numbers)
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Solve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStandardDev(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(1)
	output := math.Round(utilities.StandardDev(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}

func TestVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := float64(2)
	output := math.Round(utilities.Variance(input))

	if output != expected {
		t.Fatalf("Expected:%v But got:%v", expected, output)
	}

}

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
