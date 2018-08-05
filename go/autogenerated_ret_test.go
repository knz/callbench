// auto-generated with gen.sh
package main

import "testing"

//go:noinline
func f1r(i int) (int) { return i }
//go:noinline
func f1rmulti(d,i int) (int) {
 if d == 1 { return i }
 return f1rmulti(d-1, i)
}
//go:noinline
func f2r(i int) (int,int) { return i,i }
//go:noinline
func f2rmulti(d,i int) (int,int) {
 if d == 1 { return i,i }
 return f2rmulti(d-1, i)
}
//go:noinline
func f3r(i int) (int,int,int) { return i,i,i }
//go:noinline
func f3rmulti(d,i int) (int,int,int) {
 if d == 1 { return i,i,i }
 return f3rmulti(d-1, i)
}
//go:noinline
func f4r(i int) (int,int,int,int) { return i,i,i,i }
//go:noinline
func f4rmulti(d,i int) (int,int,int,int) {
 if d == 1 { return i,i,i,i }
 return f4rmulti(d-1, i)
}
//go:noinline
func f5r(i int) (int,int,int,int,int) { return i,i,i,i,i }
//go:noinline
func f5rmulti(d,i int) (int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i }
 return f5rmulti(d-1, i)
}
//go:noinline
func f6r(i int) (int,int,int,int,int,int) { return i,i,i,i,i,i }
//go:noinline
func f6rmulti(d,i int) (int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i }
 return f6rmulti(d-1, i)
}
//go:noinline
func f7r(i int) (int,int,int,int,int,int,int) { return i,i,i,i,i,i,i }
//go:noinline
func f7rmulti(d,i int) (int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i }
 return f7rmulti(d-1, i)
}
//go:noinline
func f8r(i int) (int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i }
//go:noinline
func f8rmulti(d,i int) (int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i }
 return f8rmulti(d-1, i)
}
//go:noinline
func f9r(i int) (int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i }
//go:noinline
func f9rmulti(d,i int) (int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i }
 return f9rmulti(d-1, i)
}
//go:noinline
func f10r(i int) (int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f10rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i }
 return f10rmulti(d-1, i)
}
//go:noinline
func f11r(i int) (int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f11rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i }
 return f11rmulti(d-1, i)
}
//go:noinline
func f12r(i int) (int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f12rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i }
 return f12rmulti(d-1, i)
}
//go:noinline
func f13r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f13rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f13rmulti(d-1, i)
}
//go:noinline
func f14r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f14rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f14rmulti(d-1, i)
}
//go:noinline
func f15r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f15rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f15rmulti(d-1, i)
}
//go:noinline
func f16r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f16rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f16rmulti(d-1, i)
}
//go:noinline
func f17r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f17rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f17rmulti(d-1, i)
}
//go:noinline
func f18r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f18rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f18rmulti(d-1, i)
}
//go:noinline
func f19r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f19rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f19rmulti(d-1, i)
}
//go:noinline
func f20r(i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
//go:noinline
func f20rmulti(d,i int) (int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int,int) {
 if d == 1 { return i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i,i }
 return f20rmulti(d-1, i)
}
func BenchmarkRet1(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r := f1r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet1(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r := f1rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_1(b *testing.B) { benchMultiRet1(1, b); }
func BenchmarkMultiRet20_1(b *testing.B) { benchMultiRet1(20, b); }
func BenchmarkRet2(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _ := f2r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet2(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _ := f2rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_2(b *testing.B) { benchMultiRet2(1, b); }
func BenchmarkMultiRet20_2(b *testing.B) { benchMultiRet2(20, b); }
func BenchmarkRet3(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _ := f3r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet3(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _ := f3rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_3(b *testing.B) { benchMultiRet3(1, b); }
func BenchmarkMultiRet20_3(b *testing.B) { benchMultiRet3(20, b); }
func BenchmarkRet4(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _ := f4r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet4(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _ := f4rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_4(b *testing.B) { benchMultiRet4(1, b); }
func BenchmarkMultiRet20_4(b *testing.B) { benchMultiRet4(20, b); }
func BenchmarkRet5(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _ := f5r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet5(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _ := f5rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_5(b *testing.B) { benchMultiRet5(1, b); }
func BenchmarkMultiRet20_5(b *testing.B) { benchMultiRet5(20, b); }
func BenchmarkRet6(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _ := f6r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet6(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _ := f6rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_6(b *testing.B) { benchMultiRet6(1, b); }
func BenchmarkMultiRet20_6(b *testing.B) { benchMultiRet6(20, b); }
func BenchmarkRet7(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _ := f7r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet7(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _ := f7rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_7(b *testing.B) { benchMultiRet7(1, b); }
func BenchmarkMultiRet20_7(b *testing.B) { benchMultiRet7(20, b); }
func BenchmarkRet8(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _ := f8r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet8(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _ := f8rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_8(b *testing.B) { benchMultiRet8(1, b); }
func BenchmarkMultiRet20_8(b *testing.B) { benchMultiRet8(20, b); }
func BenchmarkRet9(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _ := f9r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet9(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _ := f9rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_9(b *testing.B) { benchMultiRet9(1, b); }
func BenchmarkMultiRet20_9(b *testing.B) { benchMultiRet9(20, b); }
func BenchmarkRet10(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _ := f10r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet10(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _ := f10rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_10(b *testing.B) { benchMultiRet10(1, b); }
func BenchmarkMultiRet20_10(b *testing.B) { benchMultiRet10(20, b); }
func BenchmarkRet11(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _ := f11r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet11(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _ := f11rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_11(b *testing.B) { benchMultiRet11(1, b); }
func BenchmarkMultiRet20_11(b *testing.B) { benchMultiRet11(20, b); }
func BenchmarkRet12(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _ := f12r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet12(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _ := f12rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_12(b *testing.B) { benchMultiRet12(1, b); }
func BenchmarkMultiRet20_12(b *testing.B) { benchMultiRet12(20, b); }
func BenchmarkRet13(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _ := f13r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet13(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _ := f13rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_13(b *testing.B) { benchMultiRet13(1, b); }
func BenchmarkMultiRet20_13(b *testing.B) { benchMultiRet13(20, b); }
func BenchmarkRet14(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _ := f14r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet14(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _ := f14rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_14(b *testing.B) { benchMultiRet14(1, b); }
func BenchmarkMultiRet20_14(b *testing.B) { benchMultiRet14(20, b); }
func BenchmarkRet15(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f15r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet15(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f15rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_15(b *testing.B) { benchMultiRet15(1, b); }
func BenchmarkMultiRet20_15(b *testing.B) { benchMultiRet15(20, b); }
func BenchmarkRet16(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f16r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet16(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f16rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_16(b *testing.B) { benchMultiRet16(1, b); }
func BenchmarkMultiRet20_16(b *testing.B) { benchMultiRet16(20, b); }
func BenchmarkRet17(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f17r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet17(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f17rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_17(b *testing.B) { benchMultiRet17(1, b); }
func BenchmarkMultiRet20_17(b *testing.B) { benchMultiRet17(20, b); }
func BenchmarkRet18(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f18r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet18(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f18rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_18(b *testing.B) { benchMultiRet18(1, b); }
func BenchmarkMultiRet20_18(b *testing.B) { benchMultiRet18(20, b); }
func BenchmarkRet19(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f19r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet19(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f19rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_19(b *testing.B) { benchMultiRet19(1, b); }
func BenchmarkMultiRet20_19(b *testing.B) { benchMultiRet19(20, b); }
func BenchmarkRet20(b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f20r(i)
  val += r
 }
 CONSUME(b, val);
}
func benchMultiRet20(d int, b *testing.B) {
 val := 1
 for i := 0; i < b.N; i++ {
  r, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := f20rmulti(d, i)
  val += r
 }
 CONSUME(b, val);
}
func BenchmarkMultiRet1_20(b *testing.B) { benchMultiRet20(1, b); }
func BenchmarkMultiRet20_20(b *testing.B) { benchMultiRet20(20, b); }
