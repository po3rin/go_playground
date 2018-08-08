package benchmark

import (
	"testing"

	"github.com/golang/snappy"
)

func BenchmarkSnappy(b *testing.B) {
	const size = 512
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
