package builder_test

import (
	"strings"
	"testing"
)

func joinWithPlus(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

func joinWithBuilder(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func joinWithBuilderAndGrow(strs ...string) string {
	var sb strings.Builder
	sb.Grow(30)
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func joinWithCapByte(strs ...string) string {
	var m = make([]byte, 0, 30)
	for _, v := range strs {
		m = append(m, v...)
	}
	return string(m)
}

func BenchmarkPlus(b *testing.B) {
	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = joinWithPlus(strs...)
	}
}

func BenchmarkBuilder(b *testing.B) {
	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = joinWithBuilder(strs...)
	}
}

func BenchmarkBuilderAndGrow(b *testing.B) {
	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = joinWithBuilderAndGrow(strs...)
	}
}

func BenchmarkCapByteArray(b *testing.B) {

	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = joinWithCapByte(strs...)
	}
}
