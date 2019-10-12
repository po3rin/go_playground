package main_test

import (
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/blas/gonum"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/netlib/blas/netlib"
)

func makeRandVec(num int) *mat.VecDense {
	data := make([]float64, num)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return mat.NewVecDense(num, data)
}

func useOpenBlas() {
	blas64.Use(netlib.Implementation{})
}

func useGonumBlas() {
	blas64.Use(gonum.Implementation{})
}

func BenchmarkOpenBlas(b *testing.B) {
	useOpenBlas()
	v := makeRandVec(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mat.Dot(v, v)
	}
}

func BenchmarkGonumBlas(b *testing.B) {
	useGonumBlas()
	v := makeRandVec(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mat.Dot(v, v)
	}
}
