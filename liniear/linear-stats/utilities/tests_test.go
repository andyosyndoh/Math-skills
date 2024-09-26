package utilities

import (
	"testing"
)

func TestPearsonCorrelation(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "person", args: args{x: []float64{0, 1, 2, 3, 4}, y: []float64{2.0, 3.0, 5.0, 7.0, 11.0}}, want: 0.9722718241315028},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonCorrelation(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("PearsonCorrelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{name: "linear", args: args{x: []float64{0, 1, 2, 3, 4}, y: []float64{2.0, 3.0, 5.0, 7.0, 11.0}}, want: 2.2, want1: 1.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LinearRegression(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("LinearRegression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LinearRegression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
