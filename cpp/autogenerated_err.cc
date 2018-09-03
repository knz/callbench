// auto-generated with gen.sh
void BenchmarkNoErr0(B* b) { benchNoErr(0, b); }
void BenchmarkExc0(B* b) { benchExc(0, b); }
void BenchmarkErr0(B* b) { benchErr(0, b); }
void BenchmarkNoErr1(B* b) { benchNoErr(1, b); }
void BenchmarkExc1(B* b) { benchExc(1, b); }
void BenchmarkErr1(B* b) { benchErr(1, b); }
void BenchmarkNoErr2(B* b) { benchNoErr(2, b); }
void BenchmarkExc2(B* b) { benchExc(2, b); }
void BenchmarkErr2(B* b) { benchErr(2, b); }
void BenchmarkNoErr3(B* b) { benchNoErr(3, b); }
void BenchmarkExc3(B* b) { benchExc(3, b); }
void BenchmarkErr3(B* b) { benchErr(3, b); }
void BenchmarkNoErr4(B* b) { benchNoErr(4, b); }
void BenchmarkExc4(B* b) { benchExc(4, b); }
void BenchmarkErr4(B* b) { benchErr(4, b); }
void BenchmarkNoErr5(B* b) { benchNoErr(5, b); }
void BenchmarkExc5(B* b) { benchExc(5, b); }
void BenchmarkErr5(B* b) { benchErr(5, b); }
void BenchmarkNoErr6(B* b) { benchNoErr(6, b); }
void BenchmarkExc6(B* b) { benchExc(6, b); }
void BenchmarkErr6(B* b) { benchErr(6, b); }
void BenchmarkNoErr7(B* b) { benchNoErr(7, b); }
void BenchmarkExc7(B* b) { benchExc(7, b); }
void BenchmarkErr7(B* b) { benchErr(7, b); }
void BenchmarkNoErr8(B* b) { benchNoErr(8, b); }
void BenchmarkExc8(B* b) { benchExc(8, b); }
void BenchmarkErr8(B* b) { benchErr(8, b); }
void BenchmarkNoErr9(B* b) { benchNoErr(9, b); }
void BenchmarkExc9(B* b) { benchExc(9, b); }
void BenchmarkErr9(B* b) { benchErr(9, b); }
void BenchmarkNoErr10(B* b) { benchNoErr(10, b); }
void BenchmarkExc10(B* b) { benchExc(10, b); }
void BenchmarkErr10(B* b) { benchErr(10, b); }
void BenchmarkNoErr20(B* b) { benchNoErr(20, b); }
void BenchmarkExc20(B* b) { benchExc(20, b); }
void BenchmarkErr20(B* b) { benchErr(20, b); }
void BenchmarkNoErr50(B* b) { benchNoErr(50, b); }
void BenchmarkExc50(B* b) { benchExc(50, b); }
void BenchmarkErr50(B* b) { benchErr(50, b); }
void BenchmarkNoErr100(B* b) { benchNoErr(100, b); }
void BenchmarkExc100(B* b) { benchExc(100, b); }
void BenchmarkErr100(B* b) { benchErr(100, b); }
void BenchmarkNoErr200(B* b) { benchNoErr(200, b); }
void BenchmarkExc200(B* b) { benchExc(200, b); }
void BenchmarkErr200(B* b) { benchErr(200, b); }
void BenchmarkNoErr500(B* b) { benchNoErr(500, b); }
void BenchmarkExc500(B* b) { benchExc(500, b); }
void BenchmarkErr500(B* b) { benchErr(500, b); }
void BenchmarkNoErr1000(B* b) { benchNoErr(1000, b); }
void BenchmarkExc1000(B* b) { benchExc(1000, b); }
void BenchmarkErr1000(B* b) { benchErr(1000, b); }
int main() {
AutoBench(
Bench(BenchmarkNoErr0),
Bench(BenchmarkExc0),
Bench(BenchmarkErr0),
Bench(BenchmarkNoErr1),
Bench(BenchmarkExc1),
Bench(BenchmarkErr1),
Bench(BenchmarkNoErr2),
Bench(BenchmarkExc2),
Bench(BenchmarkErr2),
Bench(BenchmarkNoErr3),
Bench(BenchmarkExc3),
Bench(BenchmarkErr3),
Bench(BenchmarkNoErr4),
Bench(BenchmarkExc4),
Bench(BenchmarkErr4),
Bench(BenchmarkNoErr5),
Bench(BenchmarkExc5),
Bench(BenchmarkErr5),
Bench(BenchmarkNoErr6),
Bench(BenchmarkExc6),
Bench(BenchmarkErr6),
Bench(BenchmarkNoErr7),
Bench(BenchmarkExc7),
Bench(BenchmarkErr7),
Bench(BenchmarkNoErr8),
Bench(BenchmarkExc8),
Bench(BenchmarkErr8),
Bench(BenchmarkNoErr9),
Bench(BenchmarkExc9),
Bench(BenchmarkErr9),
Bench(BenchmarkNoErr10),
Bench(BenchmarkExc10),
Bench(BenchmarkErr10),
Bench(BenchmarkNoErr20),
Bench(BenchmarkExc20),
Bench(BenchmarkErr20),
Bench(BenchmarkNoErr50),
Bench(BenchmarkExc50),
Bench(BenchmarkErr50),
Bench(BenchmarkNoErr100),
Bench(BenchmarkExc100),
Bench(BenchmarkErr100),
Bench(BenchmarkNoErr200),
Bench(BenchmarkExc200),
Bench(BenchmarkErr200),
Bench(BenchmarkNoErr500),
Bench(BenchmarkExc500),
Bench(BenchmarkErr500),
Bench(BenchmarkNoErr1000),
Bench(BenchmarkExc1000),
Bench(BenchmarkErr1000),
);
}
