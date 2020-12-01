package main

import (
	"errors"
	"testing"
)

//go:noinline
func id(arg int) int { return arg }

//go:noinline
func leafNoErr(arg int) int {
	if arg == 0 {
		// Unlikely error case.
		return -1
	}
	// Common non-error case.
	return id(arg)
}

//go:noinline
func intermediateNoErr(work int) int {
	return leafNoErr(work)
}

//go:noinline
func workNoErr(work int) int {
	var n int
	for i := 0; i < work; i++ {
		n += intermediateNoErr(work + 1)
	}
	return n
}

func benchNoErr(work int, b *testing.B) {
	val := 1
	for i := 0; i < b.N; i++ {
		val += workNoErr(work)
	}
	CONSUME(b, val)
}

///////////////////////////////////////////////////

var errObj = errors.New("woo")

//go:noinline
func leafExc(arg int) int {
	if arg == 0 {
		// Unlikely error case.
		panic(errObj)
	}
	// Common non-error case.
	return id(arg)
}

//go:noinline
func intermediateExc(arg int) int {
	return leafExc(arg)
}

//go:noinline
func workExc(work int) (res int) {
	defer func() {
		if r := recover(); r != nil {
			res = 42
		}
	}()
	var n int
	for i := 0; i < work; i++ {
		n += intermediateExc(work + 1)
	}
	return n
}

func benchExc(work int, b *testing.B) {
	val := 1
	for i := 0; i < b.N; i++ {
		val += workExc(work)
	}
	CONSUME(b, val)
}

///////////////////////////////////////////////////

//go:noinline
func leafErr(arg int) (int, error) {
	if arg == 0 {
		// Unlikely error case.
		return 0, errObj
	}
	// Common non-error case.
	return id(arg), nil
}

//go:noinline
func intermediateErr(work int) (int, error) {
	return leafErr(work)
}

//go:noinline
func workErr(work int) int {
	var n int
	for i := 0; i < work; i++ {
		val, err := intermediateErr(work)
		if err != nil {
			return 42
		}
		n += val
	}
	return n
}

func benchErr(work int, b *testing.B) {
	val := 1
	for i := 0; i < b.N; i++ {
		val += workErr(work)
	}
	CONSUME(b, val)
}
