package subx

import (
	"regexp"
	"testing"
	"time"

	c "constraints"
)

// TimeFunc describe a time value initializer with namespaced short-hand methods
// for direct test definitions.
type TimeFunc func() time.Time

// Time returns a TimeFunc that return v. The returned type is equivalent to
// that returned from Value(v), except it expose useful short-hand methods for
// defining tests directly.
func Time(v time.Time) TimeFunc {
	return func() time.Time {
		return v
	}
}

// Before is a short-hand for Test(vf, TimeBefore(w)).
func (vf TimeFunc) Before(w time.Time) func(*testing.T) {
	return Test(vf, TimeBefore(w))
}

// NotBefore is a short-hand for Test(vf, TimeNotBefore(w)).
func (vf TimeFunc) NotBefore(w time.Time) func(*testing.T) {
	return Test(vf, TimeNotBefore(w))
}

// Equal is a short-hand for Test(vf, TimeEqual(w)).
func (vf TimeFunc) Equal(w time.Time) func(*testing.T) {
	return Test(vf, TimeEqual(w))
}

// NotEqual is a short-hand for Test(vf, TimeNotEqual(w)).
func (vf TimeFunc) NotEqual(w time.Time) func(*testing.T) {
	return Test(vf, TimeNotEqual(w))
}

// IsZero is a short-hand for Test(vf, TimeIsZero()).
func (vf TimeFunc) IsZero() func(*testing.T) {
	return Test(vf, TimeIsZero())
}

// IsNotZero is a short-hand for Test(vf, TimeIsNotZero()).
func (vf TimeFunc) IsNotZero() func(*testing.T) {
	return Test(vf, TimeIsNotZero())
}

// OrderedFunc describe a ordered value initializer with namespaced short-hand
// methods for direct test definitions.
type OrderedFunc[T c.Ordered] func() T

// Ordered returns a NumericFunc that return v. The returned type is equivalent
// to that returned from Value(v), except it expose useful short-hand methods
// for defining tests directly.
func Ordered[T c.Ordered](v T) OrderedFunc[T] {
	return func() T {
		return v
	}
}

// Equal is a short-hand for Test(vf, CompareEqual(w)).
func (vf OrderedFunc[T]) Equal(w T) func(t *testing.T) {
	return Test(vf, CompareEqual(w))
}

// NotEqual is a short-hand for Test(vf, CompareNotEqual(w)).
func (vf OrderedFunc[T]) NotEqual(w T) func(t *testing.T) {
	return Test(vf, CompareNotEqual(w))
}

// GreaterOrEqual is a short-hand for Test(vf, OrderGreaterOrEqual(w)).
func (vf OrderedFunc[T]) GreaterOrEqual(w T) func(t *testing.T) {
	return Test(vf, OrderGreaterOrEqual(w))
}

// Greater is a short-hand for Test(vf, OrderGreater(w)).
func (vf OrderedFunc[T]) Greater(w T) func(t *testing.T) {
	return Test(vf, OrderGreater(w))
}

// LessOrEqual is a short-hand for Test(vf, OrderLessOrEqual(w)).
func (vf OrderedFunc[T]) LessOrEqual(w T) func(t *testing.T) {
	return Test(vf, OrderLessOrEqual(w))
}

// Less is a short-hand for Test(vf, OrderLess(w)).
func (vf OrderedFunc[T]) Less(w T) func(t *testing.T) {
	return Test(vf, OrderLess(w))
}

// StringFunc describe a string value initializer with namespaced short-hand
// methods for direct test definitions.
type StringFunc func() string

// String returns a StringFunc that return v. The returned type is equivalent
// to that returned from Value(v), except it expose useful short-hand methods
// for defining tests directly.
func String(v string) StringFunc {
	return func() string {
		return v
	}
}

// Equal is a short-hand for Test(vf, CompareEqual(w)).
func (vf StringFunc) Equal(w string) func(t *testing.T) {
	return Test(vf, CompareEqual(w))
}

// NotEqual is a short-hand for Test(vf, CompareNotEqual(w)).
func (vf StringFunc) NotEqual(w string) func(t *testing.T) {
	return Test(vf, CompareNotEqual(w))
}

// GreaterOrEqual is a short-hand for Test(vf, OrderGreaterOrEqual(w)).
func (vf StringFunc) GreaterOrEqual(w string) func(t *testing.T) {
	return Test(vf, OrderGreaterOrEqual(w))
}

// Greater is a short-hand for Test(vf, OrderGreater(w)).
func (vf StringFunc) Greater(w string) func(t *testing.T) {
	return Test(vf, OrderGreater(w))
}

// LessOrEqual is a short-hand for Test(vf, OrderLessOrEqual(w)).
func (vf StringFunc) LessOrEqual(w string) func(t *testing.T) {
	return Test(vf, OrderLessOrEqual(w))
}

// Less is a short-hand for Test(vf, OrderLess(w)).
func (vf StringFunc) Less(w string) func(t *testing.T) {
	return Test(vf, OrderLess(w))
}

// MatchRegexp is a short-hand for Test(vf, StringMatchRegexp(w)).
func (vf StringFunc) MatchRegexp(w *regexp.Regexp) func(t *testing.T) {
	return Test(vf, StringMatchRegexp(w))
}

// Contains is a short-hand for Test(vf, StringContains(w)).
func (vf StringFunc) Contains(w string) func(t *testing.T) {
	return Test(vf, StringContains(w))
}

// ContainsAny is a short-hand for Test(vf, StringContainsAny(w)).
func (vf StringFunc) ContainsAny(w string) func(t *testing.T) {
	return Test(vf, StringContainsAny(w))
}

// EqualFold is a short-hand for Test(vf, StringEqualFold(w)).
func (vf StringFunc) EqualFold(w string) func(t *testing.T) {
	return Test(vf, StringEqualFold(w))
}

// HasPrefix is a short-hand for Test(vf, StringHasPrefix(w)).
func (vf StringFunc) HasPrefix(w string) func(t *testing.T) {
	return Test(vf, StringHasPrefix(w))
}

// HasSuffix is a short-hand for Test(vf, StringHasSuffix(w)).
func (vf StringFunc) HasSuffix(w string) func(t *testing.T) {
	return Test(vf, StringHasSuffix(w))
}
