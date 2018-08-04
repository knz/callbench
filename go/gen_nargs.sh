#!/bin/sh

echo "// auto-generated with gen.sh"
echo "package main"
echo
echo 'import "testing"'
echo
for ((i=1; i<=20; i++)); do
	echo "//go:noinline"
	echo -n "func f$i("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "a$j"
	done
	echo -n " int) int { return "
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n "+"; fi
		echo -n "a$j"
	done
	echo " }"
done

for ((i=0; i<=20; i++)); do
	echo "func BenchmarkArg$i(b *testing.B) {"
	echo " val := 1"
	echo " for i := 0; i < b.N; i++ {"
	echo -n "  val += f$i("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo ");"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"

	echo "func BenchmarkVarArg$i(b *testing.B) {"
	echo " val := 1"
	echo " for i := 0; i < b.N; i++ {"
	echo -n "  val += fvar($i"
	for ((j=0;j<i;j++)); do
		echo -n ",i"
	done
	echo ");"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
done

