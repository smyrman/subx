package subx

import (
	"fmt"
)

// AnyOf is a short-and for NOf(1, checks...). This means that if len(checks) is
// 0, then the returned check will always fail.
func AnyOf[T any](checks ...CheckFunc[T]) CheckFunc[T] {
	return AtLeastNOf(1, checks...)
}

// AllOf is a short-and for NOf(len(checks), checks...). This means that if
// len(checks) is 0, then the returned check will always pass.
func AllOf[T any](checks ...CheckFunc[T]) CheckFunc[T] {
	return AtLeastNOf(len(checks), checks...)
}

// AtLeastNOf return a check that passes if at least N of the provided checks
// pass.
func AtLeastNOf[T any](n int, checks ...CheckFunc[T]) CheckFunc[T] {
	// Compose best error message prefix based on n and number of checks.
	var want string
	switch n {
	case 0:
		// Gives better performance.
		return func(vf func() T) error { return nil }
	case len(checks):
		if n == 1 {
			// Gives better error message to return first check.
			return checks[0]
		}
		want = "all checks to pass"
	default:
		want = fmt.Sprintf("at least %d/%d checks to pass", n, len(checks))
	}

	return func(vf func() T) error {
		var errs Errors
		for _, cf := range checks {
			if err := cf(vf); err != nil {
				errs = append(errs, err)
			}
		}
		if cnt := len(checks) - len(errs); cnt < n {
			return fmt.Errorf("%d/%d checks failed, want %s;\n%w", len(errs), len(checks), want, errs)
		}
		return nil
	}
}

// Repeat return a slice of size n where the passed in check is repeated N
// times. Can e.g. be useful in combination with NOf or AllOf to test if a
// function return stable results if called multiple times.
func Repeat[T any](n int, cf CheckFunc[T]) []CheckFunc[T] {
	checks := make([]CheckFunc[T], 0, n)
	for i := 0; i < n; i++ {
		checks = append(checks, cf)
	}
	return checks
}
