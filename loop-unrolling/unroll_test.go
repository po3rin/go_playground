package unroll_test

import (
	"testing"

	unroll "github.com/po3rin/unrolling"
)

func makeBenchData(n int) []uint64 {
	ids := make([]uint64, n)
	for i := 0; i < n; i++ {
		ids[i] = uint64(i)
	}
	return ids
}

var dataRen = 10000
var id uint64 = 5000

func BenchmarkContainsUint64(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64(ids, id)
	}
}

func BenchmarkContainsUint64WithMap(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64WithMap(ids, id)
	}
}

func BenchmarkContainsUint64Unroll2(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64Unroll2(ids, id)
	}
}

func BenchmarkContainsUint64Unroll4(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64Unroll2(ids, id)
	}
}

func BenchmarkContainsUint64Unroll8(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64Unroll2(ids, id)
	}
}

func BenchmarkContainsUint64Unroll8WithBoundsCheck(b *testing.B) {
	ids := makeBenchData(dataRen)
	for i := 0; i < b.N; i++ {
		unroll.ContainsUint64Unroll8WithBoundsCheck(ids, id)
	}
}
