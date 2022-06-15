// Package assert enables approx equality assertion of structs that contain floats or decimals.
package assert

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// EqualApprox evaluates if float and decimal fields of T are within margin.
func EqualApprox[T any](t *testing.T, exp, act T, margin float64) {
	equal := cmp.Equal(exp, act,
		cmp.Transformer("DecToFloat64",
			func(x decimal.Decimal) float64 {
				return x.InexactFloat64()
			}),
		cmpopts.EquateApprox(0, margin))

	assert.Truef(t, equal, "Expected:\n%s\n Actual:\n%s\n", spew.Sdump(exp), spew.Sdump(act))
}
