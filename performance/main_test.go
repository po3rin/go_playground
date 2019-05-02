package main

import (
	"testing"
)

func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // run the Fib function b.N times
	}
}

func BenchmarkFib1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(1) // run the Fib function b.N times
	}
}

func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcnt(uint64(i))
	}
}

var Result uint64

func BenchmarkPopcnt2(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = Popcnt(uint64(i))
	}
	Result = r
}

func BenchmarkPopcnt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = Popcnt(uint64(i))
	}
}

func BenchmarkPopcnt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Popcnt(uint64(i))
	}
}
