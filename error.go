package subx

import (
	"fmt"
	"bytes"
)

type cmpError struct{
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
