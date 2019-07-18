package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	u := mat.NewVecDense(3, []float64{1, 2, 3})
	println("u: ")
	matPrint(u)

	v := mat.NewVecDense(3, []float64{4, 5, 6})
	println("v: ")
	matPrint(v)

	w := mat.NewVecDense(3, nil)
	w.AddVec(u, v)
	println("u + v: ")
	matPrint(w)

	// Or, you can overwrite u with the result of your operation to
	//save space.
	u.AddVec(u, v)
	println("u (overwritten):")
	matPrint(u)

	// Add u + alpha * v for some scalar alpha
	w.AddScaledVec(u, 2, v)
	println("u + 2 * v: ")
	matPrint(w)

	// Subtract v from u
	w.SubVec(u, v)
	println("v - u: ")
	matPrint(w)

	// Scale u by alpha
	w.ScaleVec(23, u)
	println("u * 23: ")
	matPrint(w)

	// Compute the dot product of u and v
	// Since float64’s don’t have a dot method, this is not done
	//inplace
	d := mat.Dot(u, v)
	println("u dot v: ", d)

	// Find length of v
	l := v.Len()
	println("Length of v: ", l)

	// We can also find the length of v in Euclidean space
	// The 2 parameter specifices that this is the Frobenius norm
	// Rather than the maximum absolute column sum
	// This distinction is more important when Norm is applied to
	// Matrices since vectors only have one column and the the
	// maximum absolute column sum is the Frobenius norm squared.
	matPrint(v)
	println(mat.Norm(v, 2))

	// Create a new matrix
	// x := make([]float64, 12)
	// for i := 0; i < 12; i++ {
	// 	x[i] = float64(i)
	// }
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	A := mat.NewDense(3, 4, x)
	matPrint(A)

	sumOfIndices := func(i, j int, v float64) float64 {
		return float64(i+j) + v
	}
	var Ap mat.Dense
	Ap.Apply(sumOfIndices, A)
	println("apply")
	matPrint(&Ap)

	E := A.Slice(0, 3, 0, 3)
	matPrint(E)
	de := mat.Det(E)
	println("det(E): ", de)

	var B mat.Dense
	A = mat.NewDense(2, 2, []float64{3, 5, 1, 2})
	_ = B.Inverse(A)
	println("A-1:")
	matPrint(&B)

	D := Add(A, A)
	matPrint(D)

	// // var B mat.Dense
	// A.Add(A, A)
	// println("A + A:")
	// matPrint(A)

	// // var C mat.Dense
	// A.Sub(A, A)
	// println("A - A:")
	// matPrint(A)

	// println("Row 1 of A:")
	// matPrint(A.RowView(1))

	// println("Column 0 of A:")
	// matPrint(A.ColView(0))

	// row := []float64{10, 9, 8, 7}
	// A.SetRow(0, row)
	// println("Updated A:")
	// matPrint(A)

	// col := []float64{3, 2, 1}
	// A.SetCol(0, col)
	// println("Updated A:")
	// matPrint(A)

	// Setting and getting are pretty simple
	// a := A.At(0, 2)
	// println("A[0, 2]: ", a)
	// A.Set(0, 2, -1.5)
	// matPrint(A)

	// var C mat.Dense
	// C.Scale(2, A)
	// println("2 * A:")
	// matPrint(&C)

	// B := A.T()
	// matPrint(B)

	var C mat.Dense
	C.Product(A, A.T())
	println("A * A: ")
	matPrint(&C)
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
