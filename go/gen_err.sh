#!/bin/sh

echo "// auto-generated with gen.sh"
echo "package main"
echo
echo 'import "testing"'
echo
for i in `seq 0 10` 20 50 100 200 500 1000; do
	echo "func BenchmarkNoErr$i(b *testing.B) { benchNoErr($i, b); }"
	echo "func BenchmarkExc$i(b *testing.B) { benchExc($i, b); }"
	echo "func BenchmarkErr$i(b *testing.B) { benchErr($i, b); }"
done
