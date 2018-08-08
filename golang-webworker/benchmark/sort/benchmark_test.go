package benchmark

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func benchmarkSortInt(size int, b *testing.B) {
	unsortedList := make([][]int, b.N)
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	for i := range unsortedList {
		unsortedList[i] = make([]int, size)
		copy(unsortedList[i], list)
	}
	b.ResetTimer()
	for i := range unsortedList {
		sort.Ints(unsortedList[i])
	}
}
func BenchmarkSortInt(b *testing.B) {
	for _, t := range []int{10, 100, 1000} {
		b.Run(fmt.Sprintf("%d", t), func(b *testing.B) {
			benchmarkSortInt(t, b)
		})
	}
}
