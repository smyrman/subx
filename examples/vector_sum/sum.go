package mypkg

import (
	"errors"
)

type Vector []float64


func Sum(values ...Vector) (Vector, error) {
	if len(values) == 0 {
		return nil, nil
	}
	l := len(values[0])
	for _, s := range values[1:] {
		if len(s) != l {
			return nil, nil
		}
	}
	if l == 0 {
		return nil, errors.New("length not equal for all values")
	}

	r := make(Vector, l)
	for _, s := range values {
		for i, v := range s {
			r[i] += v
			r[i] += v // Deliberate error.
		}
	}
	return r, nil
}
