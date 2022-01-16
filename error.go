package subx

import (
	"bytes"
	"fmt"
)

type cmpError struct {
	suffix string
	result interface{}
	expect []interface{}
}

// Expect returns an error for when a comparison fails.
func Expect(result interface{}, rel string, expect ...interface{}) error {
	return cmpError{
		suffix: rel,
		result: result,
		expect: expect,
	}
}

func (err cmpError) Error() string {
	switch len(err.expect) {
	case 0:
		return fmt.Sprintf("comparison failed:\ngot: %s\nwant %s.",
			formatType(err.result),
			err.suffix,
		)
	case 1:
		return fmt.Sprintf("comparison failed:\ngot: %s\nwant %s: %s",
			formatType(err.result),
			err.suffix,
			formatType(err.expect[0]),
		)
	default:
		return fmt.Sprintf("comparison failed:\ngot: %s\nwant %s: %s",
			formatType(err.result),
			err.suffix,
			formatType(err.expect),
		)
	}
}

// Errors indents and combine the output of multiple errors on separate lines.
type Errors []error

func (errs Errors) Error() string {
	var buf bytes.Buffer
	for i, err := range errs {
		if err == nil {
			fmt.Fprintf(&buf, "[%d]: (nil)\n", i)
		} else {
			fmt.Fprintf(&buf, "[%d]: %s\n", i, indentStringExcludeFirstLine(err.Error()))
		}
	}
	return buf.String()
}

// Accumulator allows accumulating errors.
type Accumulator struct {
	prefix string
	errs   Errors
}

// SetPrefix to use on all errors.
func (acc *Accumulator) SetPrefix(prefix string) {
	acc.prefix = prefix
}

// Register adds an error to acc if it's not nil.
func (acc *Accumulator) Register(err error) {
	if err != nil {
		acc.errs = append(acc.errs, err)
	}
}

// Registerf adds an error to acc if it's not nil using the passed in formatting
// string.
func (acc *Accumulator) Registerf(format string, err error) {
	if err != nil {
		acc.errs = append(acc.errs, fmt.Errorf(format, err))
	}
}

// Result returns the accumulated error or nil.
func (acc *Accumulator) Result() error {
	switch len(acc.errs) {
	case 0:
		return nil
	case 1:
		if acc.prefix == "" {
			return acc.errs[0]
		}
		return fmt.Errorf("%s%w", acc.prefix, acc.errs[0])
	default:
		return fmt.Errorf("%s\n%w", acc.prefix, acc.errs)
	}
}
