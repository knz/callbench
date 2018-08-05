#!/bin/sh

echo "// auto-generated with gen.sh"
echo "package main"
echo
echo 'import "testing"'
echo
for ((i=1; i<=20; i++)); do
	echo "//go:noinline"
	echo -n "func f${i}r(i int) ("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "int"
	done
	echo -n ") { return "
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo " }"
	echo "//go:noinline"
	echo -n "func f${i}rmulti(d,i int) ("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "int"
	done
	echo ") {"
	echo -n " if d == 1 { return "
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo " }"
	echo " return f${i}rmulti(d-1, i)"
	echo "}"
done

for ((i=1; i<=20; i++)); do
	echo "func BenchmarkRet$i(b *testing.B) {"
	echo " val := 1"
	echo " for i := 0; i < b.N; i++ {"
	echo -n "  r"
	for ((j=1;j<i;j++)); do echo -n ", _"; done
	echo " := f${i}r(i)"
	echo "  val += r"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
	echo "func benchMultiRet$i(d int, b *testing.B) {"
	echo " val := 1"
	echo " for i := 0; i < b.N; i++ {"
	echo -n "  r"
	for ((j=1;j<i;j++)); do echo -n ", _"; done
	echo " := f${i}rmulti(d, i)"
	echo "  val += r"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
	echo "func BenchmarkMultiRet1_$i(b *testing.B) { benchMultiRet$i(1, b); }"
	echo "func BenchmarkMultiRet20_$i(b *testing.B) { benchMultiRet$i(20, b); }"
done

