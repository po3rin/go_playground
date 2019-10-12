package matrix_test

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

var t mat.Matrix = mat.NewDense(1000, 1000, make([]float64, 1000000, 1000000))

func BenchmarkDenseCopyFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := mat.DenseCopyOf(t)
		_ = d
	}
}

func BenchmarkAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d, ok := t.(*mat.Dense)
		if !ok {
			panic("gonnp: failed to transpose matrix to dense")
		}
		_ = d
	}
}

func BenchmarkCloneFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d mat.Dense
		d.CloneFrom(t)
		_ = d
	}
}
