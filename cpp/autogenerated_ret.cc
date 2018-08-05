// auto-generated with gen.sh
#include <tuple>
__attribute__((noinline)) std::tuple<long> f1r(long i) { return {i}; }
__attribute__((noinline)) std::tuple<long,long> f2r(long i) { return {i,i}; }
__attribute__((noinline)) std::tuple<long,long,long> f3r(long i) { return {i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long> f4r(long i) { return {i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long> f5r(long i) { return {i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long> f6r(long i) { return {i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long> f7r(long i) { return {i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long> f8r(long i) { return {i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long> f9r(long i) { return {i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long> f10r(long i) { return {i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long> f11r(long i) { return {i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long> f12r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long> f13r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long> f14r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f15r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f16r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f17r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f18r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f19r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
__attribute__((noinline)) std::tuple<long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long,long> f20r(long i) { return {i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i}; }
void BenchmarkRet1(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f1r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet2(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f2r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet3(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f3r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet4(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f4r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet5(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f5r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet6(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f6r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet7(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f7r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet8(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f8r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet9(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f9r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet10(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f10r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet11(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f11r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet12(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f12r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet13(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f13r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet14(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f14r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet15(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f15r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet16(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f16r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet17(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f17r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet18(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f18r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet19(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f19r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
void BenchmarkRet20(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  auto t = f20r(i);
  val += std::get<0>(t);
 }
 CONSUME(b, val);
}
int main() {
AutoBench(
Bench(BenchmarkRet1),
Bench(BenchmarkRet2),
Bench(BenchmarkRet3),
Bench(BenchmarkRet4),
Bench(BenchmarkRet5),
Bench(BenchmarkRet6),
Bench(BenchmarkRet7),
Bench(BenchmarkRet8),
Bench(BenchmarkRet9),
Bench(BenchmarkRet10),
Bench(BenchmarkRet11),
Bench(BenchmarkRet12),
Bench(BenchmarkRet13),
Bench(BenchmarkRet14),
Bench(BenchmarkRet15),
Bench(BenchmarkRet16),
Bench(BenchmarkRet17),
Bench(BenchmarkRet18),
Bench(BenchmarkRet19),
Bench(BenchmarkRet20),
);
}
