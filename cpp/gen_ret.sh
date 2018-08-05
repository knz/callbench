#!/bin/sh

echo "// auto-generated with gen.sh"

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
	echo -n "__attribute__((noinline)) std::tuple<"
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "long"
	done
	echo  "> f${i}rmulti(long d, long i) {"
	echo " NOP(d);"
	echo " if (d == 1) "
	echo -n "  return {"
	for ((j=0;j<i;j++)); do
		if test $j -gt 0; then echo -n ","; fi
		echo -n "i"
	done
	echo "};"
	echo " return f${i}rmulti(d-1, i);"
	echo "}"
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
	echo "void benchMultiRet$i(long d, B* b) {"
	echo " long val = 1, N = b->N;"
	echo " for (long i = 0; i < N; i++) {"
	echo "  NOP(i);"
	echo "  auto t = f${i}rmulti(d, i);"
	echo "  val += std::get<0>(t);"
	echo " }"
	echo " CONSUME(b, val);"
	echo "}"
	echo "void BenchmarkMultiRet1_$i(B* b) { benchMultiRet$i(1, b); }"
	echo "void BenchmarkMultiRet20_$i(B* b) { benchMultiRet$i(20, b); }"
done

echo "int main() {"
echo "AutoBench("
for ((i=1;i<=20;i++)); do
	echo "Bench(BenchmarkRet$i),"
	echo "Bench(BenchmarkMultiRet1_$i),"
	echo "Bench(BenchmarkMultiRet20_$i),"
done
echo ");"
echo "}"
