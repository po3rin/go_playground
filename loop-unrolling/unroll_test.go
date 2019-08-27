package unroll

import (
	"fmt"
	"testing"
)

func makeBenchData(n int) []uint64 {
	ids := make([]uint64, n)
	for i := 0; i < n; i++ {
		ids[i] = uint64(i)
	}
	return ids
}

func BenchmarkAccessStructure(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		fmt.Println("=============")
		fmt.Println("")
		benchmarkAccessStructure(b, size)
		fmt.Println("")
	}
}

func benchmarkAccessStructure(b *testing.B, size int) {
	ids := makeBenchData(size)

	b.ResetTimer()

	b.Run(fmt.Sprintf("ContainsUint64_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsUint64(ids, 5000)
		}
	})

	b.Run(fmt.Sprintf("ContainsUint64WithMap_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsUint64WithMap(ids, 5000)
		}
	})

	b.Run(fmt.Sprintf("ContainsUint64Unroll2_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsUint64Unroll2(ids, 5000)
		}
	})

	b.Run(fmt.Sprintf("ContainsUint64Unroll4_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsUint64Unroll4(ids, 5000)
		}
	})

	b.Run(fmt.Sprintf("ContainsUint64Unroll8_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsUint64Unroll8(ids, 5000)
		}
	})

}

// func BenchmarkContainsUint64withMap(b *testing.B) {
// 	ids := makeBenchData(100000)
// 	for i := 0; i < b.N; i++ {
// 		ContainsUint64WithMap(ids, 5000)
// 	}
// }

// func BenchmarkContainsUint64(b *testing.B) {
// 	ids := makeBenchData(100000)
// 	for i := 0; i < b.N; i++ {
// 		ContainsUint64(ids, 5000)
// 	}
// }

// func BenchmarkContainsUint64Unroll2(b *testing.B) {
// 	ids := makeBenchData(100000)
// 	for i := 0; i < b.N; i++ {
// 		ContainsUint64Unroll2(ids, 5000)
// 	}
// }

// func BenchmarkContainsUint64Unroll4(b *testing.B) {
// 	ids := makeBenchData(100000)
// 	for i := 0; i < b.N; i++ {
// 		ContainsUint64Unroll4(ids, 5000)
// 	}
// }

// func BenchmarkContainsUint64Unroll8(b *testing.B) {
// 	ids := makeBenchData(100000)
// 	for i := 0; i < b.N; i++ {
// 		ContainsUint64Unroll8(ids, 5000)
// 	}
// }
