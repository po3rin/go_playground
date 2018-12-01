package stringer_test

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

func joinWithStringer(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func BenchmarkPlus(b *testing.B) {
	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh"}
	var str string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str = joinWithPlus(strs...)
	}
	b.Log(str)
}

func BenchmarkStringer(b *testing.B) {
	strs := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh"}
	var str string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str = joinWithStringer(strs...)
	}
	b.Log(str)
}
