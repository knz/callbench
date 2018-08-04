package main

// go:noinline
func f0() int { return 0 }

// go:noinline
func fvar(args ...int) int {
	val := 0
	for _, a := range args {
		val += a
	}
	return val
}
