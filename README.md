# SubX

The `subx` package provide a (test) matcher library for use with Go 1.18beta1 or newer, relying on type parametrization. As Go 1.18 isn't yet released, this code is naturally unsuitable for production.

## Fundamentals

This library revolve around three main concepts:

1: A "check" or `CheckFunc`, is a function that takes a value initializer function and returns a _well formatted_ error if the initialized value does not pass the check:

```go
type CheckFunc[T any] func(func() T) error
```

2: A "value initializer" or `func () T`, is any function that return a single parameter. Because the same value function can be passed to multiple checks, it is expected to return a stable result.

3: Finally, the generic `Test[T]` function converts a value initializer and a compatible check into a test function that can be passed to `t.Run`. Note that thanks to the generic design's _type constraints_, this is a type-safe operation. Thanks to the _type inference_ system, the code is still easy to read.

Examples:

```go
vf := func() int {
	return 8
}

// Valid:
t.Run("Expect correct result", subx.Test(vf, subx.OrderGreater(0))
t.Run("Expect correct result", subx.Test(vf, subx.AllOf(
	subx.OrderGreater(0),
	subx.OrderLess(10),
)))

// Invalid:
t.Run("Expect correct result", subx.Test(vf, subx.AllOf(
	subx.OrderGreater(0),
	subx.OrderLess(10),
	subx.StringHasPrefix("F"), // Compiler error; trying to use a CheckFunc[string].
)))
```

```go
vf := func() string {
	return "FOO"
}

// Valid:
t.Run("Expect correct result", subx.Test(vf, subx.AllOf(
	subx.OrderGreater("A"),
	subx.OrderLess("Z"),
	subx.StringHasPrefix("F"),
)))

// Invalid:
t.Run("Expect correct result", subx.Test(vf, subx.AllOf(
	subx.OrderGreater(0), // Compiler error; trying to use a CheckFunc[int].
	subx.OrderLess('Z'),  // Compiler error; trying to use a CheckFunc[char].
	subx.StringHasPrefix("F"),
)))
```

## Getting started with Go 1.18

Install go 1.18beta1 with tooling:

```sh
go install golang.org/dl/go1.18beta1@latest
go1.18beta1 download
go1.18beta1 install golang.org/x/tools/gopls@latest
```

If you are using VS Code, you can configure it to use Go 1.18beta1 by running "Go: Choose Go Environment" in the command menu (CTRL+P or CMD+P).

### Git clone and test subx

Next up, git clone `subx` into the into the folder of your choosing.

```sh
git clone git@github.com:smyrman/subx
```

To test that everything is working, you can now try to run one of the example tests:

```sh
$ cd  subx/examples/incorrect_sum
$ go tool go2go test
--- FAIL: TestSum (0.00s)
    --- FAIL: TestSum/[int](2,2,1) (0.00s)
        --- FAIL: TestSum/[int](2,2,1)/Expect_correct_and_positive_result (0.00s)
            sum_test.go2:16: 1/2 checks failed, want all checks to pass;
                [0]: comparison failed:
                    got: (int)
                        0
                    want equal to: (int)
                        5
    --- FAIL: TestSum/[float64](2,-3) (0.00s)
        --- FAIL: TestSum/[float64](2,-3)/Expect_correct_result (0.00s)
            sum_test.go2:29: comparison failed:
                got: (float64)
                    0
                want equal to: (float64)
                    -1
FAIL
exit status 1
FAIL	github.com/smyrman/subx/examples/incorrect_sum	0.358s
```

### Set up your own project

To create your own project using generics, you can now use go modules more or less as you normally would

```sh
mkdir <project>
cd <project>
go1.18beta1 mod init github.com/<username>/<project>
go1.18beta1 get github.com/smyrman/subx@main
```

Make sure the go.mod file says `go 1.18`.
