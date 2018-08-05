#!/bin/sh

echo "// auto-generated with gen.sh"
echo "#include <tuple>"

for ((i=1; i<=20; i++)); do
	echo -n "__attribute__((noinline)) std::tuple<"
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "long"
	done
	echo -n "> f${i}r(long i) { return {"
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo "}; }"
done

for ((i=1; i<=20; i++)); do
	echo "void BenchmarkRet$i(B* b) {"
	echo " long val = 1, N = b->N;"
	echo " for (long i = 0; i < N; i++) {"
	echo "  NOP(i);"
	echo "  auto t = f${i}r(i);"
	echo "  val += std::get<0>(t);"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
done

echo "int main() {"
echo "AutoBench("
for ((i=1;i<=20;i++)); do
	echo "Bench(BenchmarkRet$i),"
done
echo ");"
echo "}"
