package subx

import (
	"time"
)

// TimeBefore return a check that pass if the value to check is before w.
func TimeBefore(w time.Time) CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if v.Before(w) {
			return nil
		}
		return Expect(v, "before", w)
	}
}

// TimeNotBefore return a check that pass if the value to check is not before w.
func TimeNotBefore(w time.Time) CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if !v.Before(w) {
			return nil
		}
		return Expect(v, "not before", w)
	}
}

// TimeEqual return a check that pass if the value to check represent the same
// time instant as w.
func TimeEqual(w time.Time) CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if v.Equal(w) {
			return nil
		}
		return Expect(v, "same time instance", w)
	}
}

// TimeNotEqual return a check that pass if the value to check represent the
// same time instant as w.
func TimeNotEqual(w time.Time) CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if !v.Equal(w) {
			return nil
		}
		return Expect(v, "different time instance", w)
	}
}

// TimeIsZero return a check that pass if the value to check represent the zero
// time instant, January 1, year 1, 00:00:00 UTC.
func TimeIsZero() CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if v.IsZero() {
			return nil
		}
		return Expect(v, "is zero time instance")
	}
}

// TimeIsNotZero return a check that pass if the value to check represent the
// zero time instant, January 1, year 1, 00:00:00 UTC.
func TimeIsNotZero() CheckFunc[time.Time] {
	return func(vf func() time.Time) error {
		v := vf()
		if !v.IsZero() {
			return nil
		}
		return Expect(v, "is not zero time instance")
	}
}
