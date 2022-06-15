// Package dec is a set of helpers and syntax sugar for github.com/shopspring/decimal.
package dec

import (
	"github.com/shopspring/decimal"
)

type number interface {
	~int | ~int64 | ~float64
}

// New creates a new instance of decimal.Decimal from the given number.
// Int values are cast to float64 for conversion.
// Ensures equal internal representation by decimal.Decimal of the equivalent input number.
// Thus enables expected behaviour in equality funcs such as assert.Equal.
func New[T number](v T) decimal.Decimal {
	return decimal.NewFromFloat(float64(v))
}

// Between returns true if x >= lower and x <= upper.
func Between(v, lower, upper decimal.Decimal) bool {
	return v.GreaterThanOrEqual(lower) && v.LessThanOrEqual(upper)
}

// NZ (Not Zero) returns y if x is zero.
func NZ(x, y decimal.Decimal) decimal.Decimal {
	if x.IsZero() {
		return y
	}
	return x
}
