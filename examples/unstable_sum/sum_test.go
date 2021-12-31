package main

import (
	"fmt"
	"testing"

	"github.com/smyrman/subx"
)

func TestSum(t *testing.T) {
	t.Run("[int](2,2,1)", func(t *testing.T) {
		// By declaring our own value initializer, we allow the Sum function to
		// be run again before each check.
		vf := func() int {
			return Sum[int](2, 2, 1)
		}

		// We can repeat the same check multiple times with composite syntax.
		t.Run("Expect stabler results", subx.Test(vf,
			subx.AllOf(subx.Repeat(1000, subx.CompareEqual[int](5))...),
		))
	})
	t.Run("[float64](2,-3)", func(t *testing.T) {
		// By converting our value initializer function to subx.NumericFunc, we
		// get access to short-hand methods for test declaration.
		vf := subx.OrderedFunc[float64](func() float64 {
			return Sum[float64](2, -3)
		})

		// To repeat a check multiple times with the short-hand syntax, we must
		// run the check in a for-loop.
		for i := 0; i < 1000; i++ {
			t.Run(fmt.Sprintf("Run %d of 1000", i), func(t *testing.T) {
				t.Run("Expect correct result", vf.Equal(-1))
			})
		}
	})
}
