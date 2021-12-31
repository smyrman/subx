package mypkg

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestSum(t *testing.T) {
	a := Vector{1, 0, 3}
    b := Vector{0, 1, -2}
    expect := Vector{1, 1, 1}

	result, err := Sum(a, b)

    t.Run("Expect no error", subx.Test(
		subx.Value(err),
		subx.CompareEqual[error](nil),
	))
    t.Run("Expect correct sum", subx.Test(
		subx.Value(result),
		subx.DeepEqual(expect),
	))
}
