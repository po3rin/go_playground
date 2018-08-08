package benchmark

import (
	"fmt"
	"testing"

	"github.com/golang/snappy"
)

func benchmarkSnappy(size int, b *testing.B) {
	input := MakeTestText(size)
	output := make([]byte, size*10)
	// タイマーをリセット
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// b.N はオペレーションの回数
		snappy.Encode(output, input)
	}
}

func MakeTestText(size int) []byte {
	text := []byte("tech-bookfest")
	b := make([]byte, size)
	n := 0
	for n < size {
		n += copy(b[n:], text)
	}
	return b
}

func BenchmarkSnappy(b *testing.B) {
	for _, size := range []int{10, 30, 50, 100, 1000} {
		// サイズを変えてベンチマークを行う
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			benchmarkSnappy(size, b)
		})
	}
}
