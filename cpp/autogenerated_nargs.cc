// auto-generated with gen.sh
__attribute__((noinline)) long f1(
long a0) { return a0; }
__attribute__((noinline)) long f2(
long a0,long a1) { return a0+a1; }
__attribute__((noinline)) long f3(
long a0,long a1,long a2) { return a0+a1+a2; }
__attribute__((noinline)) long f4(
long a0,long a1,long a2,long a3) { return a0+a1+a2+a3; }
__attribute__((noinline)) long f5(
long a0,long a1,long a2,long a3,long a4) { return a0+a1+a2+a3+a4; }
__attribute__((noinline)) long f6(
long a0,long a1,long a2,long a3,long a4,long a5) { return a0+a1+a2+a3+a4+a5; }
__attribute__((noinline)) long f7(
long a0,long a1,long a2,long a3,long a4,long a5,long a6) { return a0+a1+a2+a3+a4+a5+a6; }
__attribute__((noinline)) long f8(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7) { return a0+a1+a2+a3+a4+a5+a6+a7; }
__attribute__((noinline)) long f9(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8) { return a0+a1+a2+a3+a4+a5+a6+a7+a8; }
__attribute__((noinline)) long f10(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9; }
__attribute__((noinline)) long f11(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10; }
__attribute__((noinline)) long f12(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11; }
__attribute__((noinline)) long f13(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12; }
__attribute__((noinline)) long f14(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13; }
__attribute__((noinline)) long f15(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14; }
__attribute__((noinline)) long f16(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14,long a15) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15; }
__attribute__((noinline)) long f17(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14,long a15,long a16) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16; }
__attribute__((noinline)) long f18(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14,long a15,long a16,long a17) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17; }
__attribute__((noinline)) long f19(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14,long a15,long a16,long a17,long a18) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17+a18; }
__attribute__((noinline)) long f20(
long a0,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14,long a15,long a16,long a17,long a18,long a19) { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17+a18+a19; }
void BenchmarkArg0(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f0();
 }
 CONSUME(b, val);
}
void BenchmarkVarArg0(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(0);
 }
 CONSUME(b, val);
}
void BenchmarkArg1(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f1(i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg1(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(1,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg2(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f2(i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg2(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(2,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg3(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f3(i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg3(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(3,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg4(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f4(i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg4(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(4,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg5(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f5(i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg5(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(5,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg6(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f6(i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg6(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(6,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg7(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f7(i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg7(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(7,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg8(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f8(i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg8(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(8,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg9(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f9(i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg9(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(9,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg10(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f10(i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg10(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(10,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg11(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f11(i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg11(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(11,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg12(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f12(i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg12(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(12,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg13(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f13(i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg13(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(13,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg14(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f14(i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg14(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(14,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg15(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f15(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg15(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(15,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg16(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f16(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg16(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(16,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg17(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f17(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg17(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(17,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg18(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f18(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg18(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(18,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg19(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f19(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg19(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(19,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkArg20(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += f20(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
void BenchmarkVarArg20(B* b) {
 long val = 1, N = b->N;
 for (long i = 0; i < N; i++) {
  NOP(i);
  val += fvar(20,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
int main() {
AutoBench(
Bench(BenchmarkArg0),
Bench(BenchmarkVarArg0),
Bench(BenchmarkArg1),
Bench(BenchmarkVarArg1),
Bench(BenchmarkArg2),
Bench(BenchmarkVarArg2),
Bench(BenchmarkArg3),
Bench(BenchmarkVarArg3),
Bench(BenchmarkArg4),
Bench(BenchmarkVarArg4),
Bench(BenchmarkArg5),
Bench(BenchmarkVarArg5),
Bench(BenchmarkArg6),
Bench(BenchmarkVarArg6),
Bench(BenchmarkArg7),
Bench(BenchmarkVarArg7),
Bench(BenchmarkArg8),
Bench(BenchmarkVarArg8),
Bench(BenchmarkArg9),
Bench(BenchmarkVarArg9),
Bench(BenchmarkArg10),
Bench(BenchmarkVarArg10),
Bench(BenchmarkArg11),
Bench(BenchmarkVarArg11),
Bench(BenchmarkArg12),
Bench(BenchmarkVarArg12),
Bench(BenchmarkArg13),
Bench(BenchmarkVarArg13),
Bench(BenchmarkArg14),
Bench(BenchmarkVarArg14),
Bench(BenchmarkArg15),
Bench(BenchmarkVarArg15),
Bench(BenchmarkArg16),
Bench(BenchmarkVarArg16),
Bench(BenchmarkArg17),
Bench(BenchmarkVarArg17),
Bench(BenchmarkArg18),
Bench(BenchmarkVarArg18),
Bench(BenchmarkArg19),
Bench(BenchmarkVarArg19),
Bench(BenchmarkArg20),
Bench(BenchmarkVarArg20),
);
}
