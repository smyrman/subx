# SubX

SubX is an experimental test library written with the go generics experiment.
As Go generics isn't yet included in the Go language, naturally this code is
unsuitable for production. However, if you want to try out Go generics, why not
write your experimental Go generics unit tests with a Go generics enabled
library?

## Getting started

To set-up the experimental Go generis tool, follow the instructions available
[here][go2go-setup].

[go2go-setup]: https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md

### Setup a GO2PATH

Once the `go2go` tool is installed, we need a place for our code to live. While
`GOPATH` in the Go world is being [deprecated][gopath-deprecated] in
favour of go modules, to use the experimental generics tool, you actually do
need to set up a `GOPATH`; or more accurately a `GO2PATH`.

Setup `GO2PATH` and git clone `subx`.

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

To test that everything is working, you can now run one of the builtin examples:

```sh
$ cd  ~/go2/src/github.com/smyrman/subx/examples/incorrect_sum
$ go tool go2go test
--- FAIL: TestSum (0.00s)
    --- FAIL: TestSum/[int](2,2,1) (0.00s)
        --- FAIL: TestSum/[int](2,2,1)/Expect_correct_and_positive_result (0.00s)
            sum_test.go2:16: 1/2 checks failed, want all checks to pass;
                [0]: subx.cmpError
                    comparison failed:
                    got: (int)
                        0
                    expect equal to: (int)
                        5
    --- FAIL: TestSum/[float64](2,-3) (0.00s)
        --- FAIL: TestSum/[float64](2,-3)/Expect_correct_result (0.00s)
            sum_test.go2:28: comparison failed:
                got: (float64)
                    0
                expect equal to: (float64)
                    -1
FAIL
exit status 1
FAIL	github.com/smyrman/subx/examples/incorrect_sum	0.362s
/Users/smyrman/Code/goroot/bin/go [test] failed: exit status 1
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
