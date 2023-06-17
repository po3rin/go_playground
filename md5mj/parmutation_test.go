package main

import (
	"testing"
)

func BenchmarkBacktrackPermutate(b *testing.B) {
	candidate := "5p1z1s7s4m7s1z3z2s2z1s"
	candidateSplit := splitN(candidate, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = backtrackPermutate(candidateSplit)
	}
}

func BenchmarkRepetitionPermute(b *testing.B) {
	candidate := "5p1z1s7s4m7s1z3z2s2z1s"
	candidateSplit := splitN(candidate, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repetitionPermute(candidateSplit)
	}
}

func BenchmarkHeapPermutate(b *testing.B) {
	candidate := "5p1z1s7s4m7s1z3z2s2z1s"
	candidateSplit := splitN(candidate, 2)
	permutations := &[][]string{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapPermutate(candidateSplit, len(candidateSplit), permutations)
	}
}

// func BenchmarkBacktrackPermutateWithDuplicate(b *testing.B) {
// 	candidate := "1m1m1m1m2p2p2p2p3s3s"
// 	candidateSplit := splitN(candidate, 2)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = backtrackPermutate(candidateSplit)
// 	}
// }

// func BenchmarkRepetitionPermuteWithDuplicate(b *testing.B) {
// 	candidate := "1m1m1m1m2p2p2p2p3s3s"
// 	candidateSplit := splitN(candidate, 2)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = repetitionPermute(candidateSplit)
// 	}
// }

// func BenchmarkHeapPermutateWithDuplicate(b *testing.B) {
// 	candidate := "1m1m1m1m2p2p2p2p3s3s"
// 	candidateSplit := splitN(candidate, 2)
// 	permutations := &[][]string{}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		heapPermutate(candidateSplit, len(candidateSplit), permutations)
// 	}
// }
