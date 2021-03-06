#include "bench.hh"
#include <tuple>
#include <stdexcept>

#define likely(x) __builtin_expect(!!(x), 1)
#define unlikely(x) __builtin_expect(!!(x), 0)
#define noinline __attribute__((noinline))

noinline
long id(long arg) { NOP(arg); return arg; }

noinline
long leafNoErr(long arg) {
	NOP(arg);
	if (unlikely(arg == 0))
		// Unlikely error case.
		return -1;
	// Common non-error case.
	return id(arg);
}

noinline
long intermediateNoErr(long arg) { NOP(arg); return leafNoErr(arg); }

noinline
long workNoErr(long work) {
	long n = 0;
	for (long i = 0; i < work; i++)
		n += intermediateNoErr(work + 1);
	return n;
}

void benchNoErr(long work, B* b) {
	long val = 1;
	for (long i = 0; i < b->N; i++) {
		NOP(i); NOP(work);
		val += workNoErr(work);
	}
	CONSUME(b, val);
}

////////////////////////////////////////////////////////

struct error{ virtual std::string what() const = 0; };
struct errImpl : error {
	std::string _reason;
	errImpl(const std::string& msg) : _reason(msg) {}
	std::string what() const { return _reason; }
};
error * errObj = new errImpl{"woo"};


noinline
long leafExc(long arg) {
	NOP(arg);
	if (unlikely(arg == 0))
		// Unlikely error case.
		throw errObj;
	// Common non-error case.
	return id(arg);
}

noinline
long intermediateExc(long arg) { NOP(arg); return leafExc(arg); }

noinline
long workExc(long work) {
	try {
		long n = 0;
		for (long i = 0; i < work; i++)
			n += intermediateExc(work + 1);
		return n;
	} catch (error *e) {
		return 42;
	}
}

void benchExc(long work, B* b) {
	long val = 1;
	for (long i = 0; i < b->N; i++) {
		NOP(i); NOP(work);
		val += workExc(work);
	}
	CONSUME(b, val);
}

/////////////////////////////////////////////////////

typedef std::tuple<long, error*> result;

noinline
result leafErr(long arg) {
	NOP(arg);
	if (unlikely(arg == 0))
		// Unlikely error case.
		return {0, errObj};
	// Common non-error case.
	return {id(arg), NULL};
}

noinline
result intermediateErr(long arg) { NOP(arg); return leafErr(arg); }

noinline
long workErr(long work) {
	long n = 0, leafWork = work + 1;
	for (long i = 0; i < work; i++) {
		NOP(i); NOP(leafWork);
		auto res = intermediateErr(leafWork);
		if (unlikely(std::get<1>(res)))
			return 42;
		n += std::get<0>(res);
	}
	return n;
}

void benchErr(long work, B* b) {
	long val = 1;
	for (long i = 0; i < b->N; i++) {
		NOP(i); NOP(work);
		val += workErr(work);
	}
	CONSUME(b, val);
}

#include "autogenerated_err.cc"
