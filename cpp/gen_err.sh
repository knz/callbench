#!/bin/sh

echo "// auto-generated with gen.sh"

for i in `seq 0 10` 20 50 100 200 500 1000; do
	echo "void BenchmarkNoErr$i(B* b) { benchNoErr($i, b); }"
	echo "void BenchmarkExc$i(B* b) { benchExc($i, b); }"
	echo "void BenchmarkErr$i(B* b) { benchErr($i, b); }"
	#for j in 0 1 2 5 10 20 50 100 1000; do
#		echo "void BenchmarkWideErr${i}_$j(B* b) { benchErrWide($j, $i, b); }"
#		echo "void BenchmarkWideExc${i}_$j(B* b) { benchExcWide($j, $i, b); }"
#	done
done

echo "int main() {"
#echo "g_benchTime = Duration(2000*1000*1000);"
echo "AutoBench("
for i in `seq 0 10` 20 50 100 200 500 1000; do
	echo "Bench(BenchmarkNoErr$i),"
	echo "Bench(BenchmarkExc$i),"
	echo "Bench(BenchmarkErr$i),"
#	for j in 0 1 2 5 10 20 50 100 1000; do
#		echo "Bench(BenchmarkWideExc${i}_$j),"
#		echo "Bench(BenchmarkWideErr${i}_$j),"
#	done
done
echo ");"
echo "}"
