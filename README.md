# callbench

Microbenchmarks to measure the costs of calling conventions in Go and C++.

Included:

- `cpp/nargs.cc`, `go/nargs_test.go`: measure the impact of adding more
  arguments to a function. Measure explict arguments vs. variable
  argument lists.
