// auto-generated with gen.sh
package main

import "testing"

//go:noinline
func f1(a0 int) int { return a0 }
//go:noinline
func f2(a0,a1 int) int { return a0+a1 }
//go:noinline
func f3(a0,a1,a2 int) int { return a0+a1+a2 }
//go:noinline
func f4(a0,a1,a2,a3 int) int { return a0+a1+a2+a3 }
//go:noinline
func f5(a0,a1,a2,a3,a4 int) int { return a0+a1+a2+a3+a4 }
//go:noinline
func f6(a0,a1,a2,a3,a4,a5 int) int { return a0+a1+a2+a3+a4+a5 }
//go:noinline
func f7(a0,a1,a2,a3,a4,a5,a6 int) int { return a0+a1+a2+a3+a4+a5+a6 }
//go:noinline
func f8(a0,a1,a2,a3,a4,a5,a6,a7 int) int { return a0+a1+a2+a3+a4+a5+a6+a7 }
//go:noinline
func f9(a0,a1,a2,a3,a4,a5,a6,a7,a8 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8 }
//go:noinline
func f10(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9 }
//go:noinline
func f11(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10 }
//go:noinline
func f12(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11 }
//go:noinline
func f13(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12 }
//go:noinline
func f14(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13 }
//go:noinline
func f15(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14 }
//go:noinline
func f16(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15 }
//go:noinline
func f17(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16 }
//go:noinline
func f18(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17 }
//go:noinline
func f19(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17,a18 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17+a18 }
//go:noinline
func f20(a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17,a18,a19 int) int { return a0+a1+a2+a3+a4+a5+a6+a7+a8+a9+a10+a11+a12+a13+a14+a15+a16+a17+a18+a19 }
func BenchmarkArg0(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f0();
 }
 CONSUME(b, val);
}
func BenchmarkVarArg0(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(0);
 }
 CONSUME(b, val);
}
func BenchmarkArg1(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f1(i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg1(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(1,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg2(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f2(i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg2(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(2,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg3(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f3(i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg3(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(3,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg4(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f4(i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg4(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(4,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg5(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f5(i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg5(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(5,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg6(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f6(i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg6(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(6,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg7(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f7(i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg7(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(7,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg8(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f8(i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg8(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(8,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg9(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f9(i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg9(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(9,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg10(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f10(i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg10(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(10,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg11(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f11(i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg11(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(11,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg12(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f12(i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg12(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(12,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg13(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f13(i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg13(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(13,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg14(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f14(i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg14(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(14,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg15(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f15(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg15(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(15,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg16(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f16(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg16(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(16,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg17(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f17(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg17(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(17,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg18(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f18(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg18(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(18,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg19(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f19(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg19(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(19,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkArg20(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += f20(i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
func BenchmarkVarArg20(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  val += fvar(20,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i);
 }
 CONSUME(b, val);
}
