package subx

import (
	"testing"
)

// Value return a value initializer that always return v. Only use if v is not
// stateful or otherwise subject to check mutation.
func Value[T any](v T) func() T {
	return func() T {
		return v
	}
}

// CheckFunc defines a function for checking a value of type T. The function is
// passed a value initializer that returns the value to check. A value
// initializer should return an equivalent value each time it's called, so that
// multiple tests can be run against the same initializer.
type CheckFunc[T any] func(func() T) error

// Test returns a test function that fail if cf(vf) returns an error. The
// returned function is marked as a helper.
func Test[T any](vf func() T, cf CheckFunc[T]) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()
		err := cf(vf)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}
