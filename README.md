# SubX

The `subx` package provide a test matcher library written with the go generics
experimental tools. As Go generics isn't yet included in the Go language,
naturally this code is unsuitable for production. However, if you want to try
out Go generics, why not conduct your experiments using a Go generics enabled
library for unit tests?

## Fundamentals

This library revolve around only three main concepts:

1: A "check" or `CheckFunc`, is a function that gets passed a value initializer
function and returns a _well formatted_ error if the returned value does not
pass the check:

```go
type CheckFunc[T any] func(func() T) error
```

2: A "value initializer" or `func () T` is any function that return a single
parameter of any type. The type`T` must match the type of the check it's being
used with. A value initializer must also accept being called multiple times as
it can be passed to multiple checks.

3: Finally, the generic `Test` function converts a value initializer and check
of the same type into a regular Go test function that can be passed to a `t.Run`
entry. Note that thanks to the generic design's _type constraints_, this is a
type-safe operation. Thanks to the _type inference_ system, the syntax is still
easy to rad.

Examples:

```go
vf := func() int {
	return 8
}

// Valid:
t.Run("Expect correct result", Test(vf, subx.OrderGreater(0))
t.Run("Expect correct result", Test(vf, subx.AllOf(
	subx.OrderGreater(0),
	subx.OrderLess(10),
)

// Invalid:
t.Run("Expect correct result", Test(vf, subx.AllOf(
	subx.OrderGreater(0),
	subx.OrderLess(10),
	subx.StringHasPrefix("F"), // Compiler error; trying to use a CheckFunc[string].
)
```

```go
vf := func() string {
	return "FOO"
}

// Valid:
t.Run("Expect correct result", Test(vf, subx.AllOf(
	subx.OrderGreater("A"),
	subx.OrderLess("Z"),
	subx.StringHasPrefix("F"),
)

// Invalid:
t.Run("Expect correct result", Test(vf, subx.AllOf(
	subx.OrderGreater(0), // Compiler error; trying to use a CheckFunc[int].
	subx.OrderLess('Z'),  // Compiler error; trying to use a CheckFunc[char].
	subx.StringHasPrefix("F"),
)
```


## Getting started

To set-up the experimental Go generis tool, follow the instructions available
[here][go2go-setup].

[go2go-setup]: https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md

### Setup a GO2PATH

Once the `go2go` tool is installed, we need a place for our code to live. While
`GOPATH` in the Go world is being [deprecated][gopath-deprecated] in
favour of go modules, to use the experimental generics tool, you actually do
need to set up a `GOPATH`; or more accurately a `GO2PATH`:

```sh
mkdir go2
export GO2PATH="$HOME/go2"
```

[gopath-deprecated]: https://blog.golang.org/go116-module-changes

### Git clone and test subx

Next up, we must git clone `subx` into the correct sub-path in GO2PATH.

```sh
mkdir -p ~/go2/src/github.com/smyrman
cd  ~/go2/src/github.com/smyrman
git clone git@github.com:smyrman/subx
```

To test that everything is working, you can now try to run one of the example
tests:

```sh
$ cd  ~/go2/src/github.com/smyrman/subx/examples/incorrect_sum
$ go tool go2go test
--- FAIL: TestSum (0.00s)
    --- FAIL: TestSum/[int](2,2,1) (0.00s)
        --- FAIL: TestSum/[int](2,2,1)/Expect_correct_and_positive_result (0.00s)
            sum_test.go2:16: 1/2 checks failed, want all checks to pass;
                [0]: comparison failed:
                    got: (int) 0
                    expect equal to: (int) 5
    --- FAIL: TestSum/[float64](2,-3) (0.00s)
        --- FAIL: TestSum/[float64](2,-3)/Expect_correct_result (0.00s)
            sum_test.go2:29: comparison failed:
                got: (float64) 0
                expect equal to: (float64) -1
FAIL
exit status 1
FAIL	github.com/smyrman/subx/examples/incorrect_sum	0.290s
```

**PS!** Note that at the time of writing, you can not specify a _path_ like you
can for `go test` when using the `go2go` tool. This will just result in an
error:

```sh
$ cd  ~/go2/src/github.com/smyrman
$ go tool go2go test ./subx/examples/incorrect_sum
package github.com/smyrman/subx/examples/incorrect_sum: no Go files in ...
```

### Set up your own project

Setup a place for your own product to live (replace `<import path>`):

```sh
mkdir -p "~/go2/src/<import path>"
cd  "~/go2/src/<import path>"
```

Create your own code using the `.go2` extension. You should now be able to
import subx as long as it's in your GO2PATH.

Example:

```go
// from examples/incorrect_sum

func TestSum(t *testing.T) {
	t.Run("[int](2,2,1)", func(t *testing.T) {
		// subx.Value runs the Sum function only once.
		vf := subx.Value(Sum[int](2, 2, 1))

		// The composite syntax can be used to combine multiple checks in one
		// sub-test.
		t.Run("Expect correct and positive result", subx.Test(vf,
		 	subx.AllOf(
				 subx.CompareEqual(5),
				 subx.OrderGreaterOrEqual(0),
			),
		))
	})
	t.Run("[float64](2,-3)", func(t *testing.T) {
		// subx.Numeric is equivalent to subx.Value, but provide short-hand
		// methods for test declaration. The Sum function is run only once.
		vf := subx.Numeric(Sum[float64](2, -3))

		// Run checks as individual tests using the short-hand methods.
		t.Run("Expect correct result", vf.Equal(-1))
		t.Run("Expect negative result", vf.LessOrEqual(0))
	})
}
```

```go
// from examples/unstable_sum

func TestSum(t *testing.T) {
	t.Run("[int](2,2,1)", func(t *testing.T) {
		// By declaring our own value initializer, we allow the Sum function to
		// be run again before each check.
		vf := func() int {
			return Sum[int](2, 2, 1)
		}

		// We can repeat the same check multiple times with composite syntax.
		t.Run("Expect stabler results", subx.Test(vf,
		 	subx.AllOf(subx.Repeat(1000, subx.CompareEqual[int](5))...),
		))
	})
	t.Run("[float64](2,-3)", func(t *testing.T) {
		// By converting our value initializer function to subx.NumericFunc, we
		// get access to short-hand methods for test declaration.
		vf := subx.NumericFunc[float64](func() float64 {
			return Sum[float64](2, -3)
		})

		// To repeat a check multiple times with the short-hand syntax, we must
		// run the check in a for-loop.
		for i := 0; i < 1000; i++ {
			t.Run(fmt.Sprintf("Run %d of 1000", i), func(t *testing.T){
				t.Run("Expect correct result", vf.Equal(-1))
			})
		}
	})
}
```

For more ideas of what to do next, read through the code to explore available
checks.
