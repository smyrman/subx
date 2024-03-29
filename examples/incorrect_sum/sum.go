package main

import (
	"constraints"
	"fmt"

	"github.com/smyrman/subx"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[T Number](values ...T) T {
	// Deliberately broken implementation.
	var v T
	for v := range values {
		v += v
	}
	return v
}

func main() {
	// Checks can be run outside of tests.
	{
		vf := func() int {
			return Sum[int](5, 1, 2)
		}
		cf := subx.CompareEqual[int](8)
		fmt.Println("CHECK Sum[int](5,1,2):", cf(vf))
	}
	{
		vf := func() float64 {
			return Sum[float64](5, 1, 2)
		}
		cf := subx.CompareEqual[float64](8)
		fmt.Println("CHECK Sum[float64](5,1,2):", cf(vf))
	}
}
