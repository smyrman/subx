package main

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestSum(t *testing.T) {
	t.Run("[int](2,2,1)", func(t *testing.T) {
		// subx.Value runs the Sum function only once.
		vf := subx.Value(Sum[int](2, 2, 1))

		// The composite syntax can be used to combine multiple checks in one
		// sub-test.
		t.Run("Expect correct and positive result", subx.Test(vf,
			subx.AllOf(
				subx.CompareEqual(5),
				subx.OrderGreaterOrEqual(0),
			),
		))
	})
	t.Run("[float64](2,-3)", func(t *testing.T) {
		// subx.Numeric is equivalent to subx.Value, but provide short-hand
		// methods for test declaration. The Sum function is run only once.
		vf := subx.Ordered(Sum[float64](2, -3))

		// Run checks as individual tests using the short-hand methods.
		t.Run("Expect correct result", vf.Equal(-1))
		t.Run("Expect negative result", vf.LessOrEqual(0))
	})
}
