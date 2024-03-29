package subx

import (
	"fmt"
	"reflect"
)

// DeepEqual returns a check function that pass when the value to check is deep
// equal to w.
func DeepEqual[T any](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if reflect.DeepEqual(v, w) {
			return nil
		}
		return Expect(v, "deep equal to", w)
	}
}

// NotDeepEqual returns a check function that pass when the value to check is
// not deep equal to w.
func NotDeepEqual[T any](w T) CheckFunc[T] {
	return func(vf func() T) error {
		v := vf()
		if !reflect.DeepEqual(v, w) {
			return nil
		}
		return Expect(v, "not deep equal to", w)
	}
}

// ReflectNil returns a check function that pass when the value to check is set
// to a typed or untyped nil value. Unlike IsNil in the reflect package, the
// check does not panic, but will error if the type can not hold a nil value.
func ReflectNil[T any]() CheckFunc[T] {
	return func(vf func() T) (err error) {
		rv := reflect.ValueOf(vf())
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("type %s of kind %s is not compatible with the ReflectNil check", rv.Type(), rv.Kind())
			}
		}()

		if rv.IsNil() {
			return nil
		}
		return Expect(rv, "reflect to nil")
	}
}

// ReflectNotNil returns a check function that pass when the value to check is
// not set to a typed or untyped nil value. Unlike IsNil in the reflect package,
// the check does not panic, but will error if the type can not hold a nil
// value.
func ReflectNotNil[T any]() CheckFunc[T] {
	return func(vf func() T) (err error) {
		rv := reflect.ValueOf(vf())
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("type %s of kind %s is not compatible with the ReflectNotNil check", rv.Type(), rv.Kind())
			}
		}()

		if !rv.IsNil() {
			return nil
		}
		return Expect(rv, "does not reflect to nil")
	}
}

// ReflectZero returns a check function that pass when the value to check is a
// zero value.
func ReflectZero[T any]() CheckFunc[T] {
	return func(vf func() T) (err error) {
		rv := reflect.ValueOf(vf())

		// rv.IsZero() can panic only if rv is invalid.
		if rv.IsZero() {
			return nil
		}
		return Expect(rv, "reflect to nil")
	}
}

// ReflectNotZero returns a check function that pass when the value to check is
// not a zero value.
func ReflectNotZero[T any]() CheckFunc[T] {
	return func(vf func() T) (err error) {
		rv := reflect.ValueOf(vf())

		// rv.IsZero() can panic only if rv is invalid.
		if !rv.IsZero() {
			return nil
		}
		return Expect(rv, "reflect to nil")
	}
}
