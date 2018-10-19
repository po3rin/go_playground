package main

import (
	"testing"
)

func BenchmarkLoop1(b *testing.B) {
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		// 都度append
		Loop1()
	}
}

func BenchmarkLoop2(b *testing.B) {
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		Loop2()
	}
}
