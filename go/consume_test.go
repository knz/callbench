package main

import "testing"

//go:noinline
func CONSUME(b *testing.B, _ ...interface{}) { b.StopTimer() }
