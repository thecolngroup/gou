// Package gou hosts simple helper functions and types for common tasks.
package gou

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Round2DP rounds x to 2 decimal places.
func Round2DP(x float64) float64 {
	return math.Round(x*100) / 100
}

// RoundTo rounds x to the nearest y precision.
func RoundTo(x, y float64) float64 {
	return y * math.Round(x/y)
}

// NN (Not Number) returns y if x is NaN or Inf.
func NN(x, y float64) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return y
	}
	return x
}

// NNZ (Not Number or Zero) returns y if x is NaN or Inf or Zero.
func NNZ(x, y float64) float64 {
	if NN(x, y) == y || x == 0 {
		return y
	}
	return x
}

// Between returns true if x >= lower and x <= upper.
func Between[T constraints.Ordered](v, lower, upper T) bool {
	return v >= lower && v <= upper
}
