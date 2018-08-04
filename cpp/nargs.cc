#include "bench.hh"
#include <cstdarg>

#define nothing __asm__ __volatile__("");

__attribute__((noinline))
long f0() {
	long ret = 0;
	NOP(ret);
	return ret;
}

__attribute__((noinline))
long fvar(size_t nargs, ...) {
	va_list ap;
	va_start(ap, nargs);
	long val = 0;
	for (size_t i = 0; i < nargs; i++) {
		val += va_arg(ap, long);
	}
	va_end(ap);
	return val;
}

#include "autogenerated_nargs.cc"
