package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {
	b := tensor.New(tensor.WithBacking(tensor.Range(tensor.Float32, 0, 24)), tensor.WithShape(2, 3, 4))
	x, _ := b.At(0, 1, 2)
	fmt.Printf("x: %v\n", x)

	// Setting data
	b.SetAt(float32(1000), 0, 1, 2)
	fmt.Printf("b:\n%v", b)
}
