package subx

// CompareEqual returns a check function that pass when the value to check
// compares equals w.
func CompareEqual[T comparable](w T) CheckFunc[T] {
	return CheckFunc[T](func(vf func() T) error {
		v := vf()
		if v == w {
			return nil
		}
		return Expect(v, "equal to", w)
	})
}

// CompareNotEqual returns a check function that pass when the value to check
// compares unequal to w.
func CompareNotEqual[T comparable](w T) CheckFunc[T] {
	return CheckFunc[T](func(vf func() T) error {
		v := vf()
		if v == w {
			return nil
		}
		return Expect(v, "unequal to", w)
	})
}
