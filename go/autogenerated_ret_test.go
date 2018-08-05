// auto-generated with gen.sh
package main

import "testing"

//go:noinline
func f1r(i int) (int) { return i }
//go:noinline
func f2r(i int) (int,int) { return i,i }
//go:noinline
func f3r(i int) (int,int,int) { return i,i,i }
//go:noinline
func f4r(i int) (int,int,int,int) { return i,i,i,i }
//go:noinline
func f5r(i int) (int,int,int,int,int) { return i,i,i,i,i }
//go:noinline
func f6r(i int) (int,int,int,int,int,int) { return i,i,i,i,i,i }
//go:noinline
func f7r(i int) (int,int,int,int,int,int,int) { return i,i,i,i,i,i,i }
//go:noinline
func f8r(i int) (int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i }
//go:noinline
func f9r(i int) (int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i }
//go:noinline
func f10r(i int) (int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f11r(i int) (int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f12r(i int) (int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f13r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f14r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f15r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f16r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f17r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f18r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f19r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f20r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
func BenchmarkRet1(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r := f1r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet2(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _ := f2r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet3(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _ := f3r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet4(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _ := f4r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet5(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _ := f5r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet6(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _ := f6r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet7(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _ := f7r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet8(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _ := f8r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet9(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _ := f9r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet10(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _ := f10r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet11(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _ := f11r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet12(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _ := f12r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet13(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _ := f13r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet14(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _ := f14r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet15(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f15r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet16(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f16r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet17(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f17r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet18(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f18r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet19(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f19r(i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkRet20(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f20r(i)
  val += r
 }
 CONSUME(b, val);
}
