package subx

import c "constraints"

func OrderGreaterOrEqual[T c.Ordered](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if v >= w {
			return nil
		}
		return Expect(v, "greater than or equal to", w)
	}
}

func OrderGreater[T c.Ordered](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if v > w {
			return nil
		}
		return Expect(v, "greater than", w)
	}
}

func OrderLessOrEqual[T c.Ordered](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if v <= w {
			return nil
		}
		return Expect(v, "less than or equal to", w)
	}
}

func OrderLess[T c.Ordered](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if v < w {
			return nil
		}
		return Expect(v, "less than", w)
	}
}
