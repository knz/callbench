#!/bin/sh

echo "// auto-generated with gen.sh"

for ((i=1; i<=20; i++)); do
	echo "__attribute__((noinline)) long f$i("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "long a$j"
	done
	echo -n ") { return "
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n "+"; fi
		echo -n "a$j"
	done
	echo "; }"
done

for ((i=0; i<=20; i++)); do
	echo "void BenchmarkArg$i(B* b) {"
	echo " long val = 1, N = b->N;"
	echo " for (long i = 0; i < N; i++) {"
	echo "  NOP(i);"
	echo -n "  val += f$i("
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo ");"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
	
	echo "void BenchmarkVarArg$i(B* b) {"
	echo " long val = 1, N = b->N;"
	echo " for (long i = 0; i < N; i++) {"
	echo "  NOP(i);"
	echo -n "  val += fvar($i"
	for ((j=0;j<i;j++)); do
		echo -n ",i"
	done
	echo ");"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
done

echo "int main() {"
echo "AutoBench("
for ((i=0;i<=20;i++)); do
	echo "Bench(BenchmarkArg$i),"
	echo "Bench(BenchmarkVarArg$i),"
done
echo ");"
echo "}"
