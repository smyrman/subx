package subx

import (
	"strings"
	"regexp"
)

func StringMatchRegexp(w *regexp.Regexp) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if w.MatchString(v) {
			return nil
		}
		return Expect(v, "match regular expression", w)
	}
}

// StringContains return a check that pass if the value to test contains the
// substring w.
func StringContains(w string) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if strings.Contains(v, w) {
			return nil
		}
		return Expect(v, "contains substring", w)
	}
}

// StringContainsAny return a check that pass if the value to test contains any
// of the Unicode code points in w.
func StringContainsAny(w string) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if strings.ContainsAny(v, w) {
			return nil
		}
		return Expect(v, "contains at least one Unicode character from", w)
	}
}

// StringEqualFold return a check that pass if the value to test interpreted as
// an UTF-8 string is equal to w under Unicode case-folding.
func StringEqualFold(w string) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if strings.EqualFold(v, w) {
			return nil
		}
		return Expect(v, "equal fold to", w)
	}
}

// StringHasPrefix return a check that pass if the value to test has the has the
// prefix w.
func StringHasPrefix(w string) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if strings.HasPrefix(v, w) {
			return nil
		}
		return Expect(v, "has prefix", w)
	}
}

// StringHasSuffix return a check that pass if the value to test has the has the
// suffix w.
func StringHasSuffix(w string) CheckFunc[string] {
	return func(vf func() string) error{
		v := vf()
		if strings.HasSuffix(v, w) {
			return nil
		}
		return Expect(v, "has suffix", w)
	}
}
