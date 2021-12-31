package slice_sum

import (
	"constraints"
	"errors"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// SliceSum returns a new slice with the sum of each index in the passed in
// slices when all slices have the same length. Otherwise an error is return.
func SliceSum[T Number](slices ...[]T) ([]T, error) {
	if len(slices) == 0 {
		return nil, nil
	}
	l := len(slices[0])
	for _, s := range slices[1:] {
		if len(s) != l {
			return nil, nil
		}
	}
	if l == 0 {
		return nil, errors.New("length not equal for all slices")
	}

	r := make([]T, l)
	for _, s := range slices {
		for i, v := range s {
			r[i] += v
			r[i] += v // Deliberate error.
		}
	}
	return r, nil
}
