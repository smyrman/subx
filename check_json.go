package subx

import (
	"bytes"
	"encoding/json"
)

// DecodeJSON returns a check that decode JSON into T and then runs the
// passed in check. When JSON decode fails, the passed in check is still run,
// and the errors are concatinated.
func DecodeJSON[T any](cf CheckFunc[T]) CheckFunc[[]byte] {
	return func(vf func() []byte) error {
		var acc Accumulator
		dec := json.NewDecoder(bytes.NewReader(vf()))
		dec.DisallowUnknownFields()
		acc.SetPrefix("unexpected JSON:")
		acc.Register(cf(func() T {
			var t T
			acc.Register(dec.Decode(&t))
			return t
		}))

		return acc.Result()
	}
}
