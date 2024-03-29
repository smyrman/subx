package main

import (
	"constraints"
	"fmt"
	"math/rand"

	"github.com/smyrman/subx"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[T Number](values ...T) T {
	var v T
	for _, _v := range values {
		v += _v
		// Deliberately broken: Subject to random error.
		// types.
		if rand.Intn(1000) == 0 {
			v--
		}
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
		cf = subx.AllOf(subx.Repeat(1000, cf)...)
		fmt.Println("CHECK Sum[int](5,1,2):", cf(vf))
	}
	{
		vf := func() float64 {
			return Sum[float64](5, 1, 2)
		}
		cf := subx.CompareEqual[float64](8)
		cf = subx.AllOf(subx.Repeat(1000, cf)...)
		fmt.Println("CHECK Sum[float64](5,1,2):", cf(vf))
	}
}
