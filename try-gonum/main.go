package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(3, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
	})

	// 各要素の値に行番号、列番号を足す
	sumOfIndices := func(i, j int, v float64) float64 {
		return float64(i+j) + v
	}

	B := mat.NewDense(3, 4, nil)
	B.Apply(sumOfIndices, A)
	matPrint(B)
	// | 1   3   5   7|
	// | 6   8  10  12|
	// |11  13  15  17|
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func Add(a mat.Matrix, b mat.Matrix) mat.Matrix {
	var B mat.Dense
	B.Add(a, b)
	return &B
}
