package slice_test // import "make_slice"

import "testing"

func BenchmarkAppend_SliceWithKey(b *testing.B) {
	stats := []string{"a", "b", "c", "d", "e", "f", "g"}
	b.ResetTimer()
	keys := make([]string, len(stats))
	for i := 0; i < b.N; i++ {
		for i, v := range stats {
			keys[i] = v
		}
	}
}

func BenchmarkSliceWithVarAppend(b *testing.B) {
	stats := []string{"a", "b", "c", "d", "e", "f", "g"}
	b.ResetTimer()
	var keys []string
	for i := 0; i < b.N; i++ {
		for _, v := range stats {
			keys = append(keys, v)
		}
	}
}

func BenchmarkAppend_SliceWithLengthAppend(b *testing.B) {
	stats := []string{"a", "b", "c", "d", "e", "f", "g"}
	b.ResetTimer()
	keys := make([]string, 0, len(stats))
	for i := 0; i < b.N; i++ {
		for _, v := range stats {
			keys = append(keys, v)
		}
	}
}
