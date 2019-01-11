package main // import try-apache-arrow""

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

type Person struct {
	Name  string
	Age   int32    // presumably, int8 should suffice.
	Langs []string // natural or programming languages.
}

func main() {
	mem := memory.NewGoAllocator()
	bld := array.NewInt32Builder(mem)
	defer bld.Release()

	// create an array with 4 values, no null
	bld.AppendValues([]int32{1, 2, 3, 4}, nil)

	arr1 := bld.NewInt32Array() // materialize the array
	defer arr1.Release()        // make sure we release memory, eventually.

	// arr1 = [1 2 3 4]
	fmt.Printf("arr1 = %v\n", arr1)

	// create an array with 5 values, 1 null
	bld.AppendValues(
		[]int32{1, 2, 3, 4, 5},
		[]bool{true, true, true, false, true},
	)

	arr2 := bld.NewInt32Array()
	defer arr2.Release()

	// arr2 = [1 2 3 (null) 5]
	fmt.Printf("arr2 = %v\n", arr2)

	sli := array.NewSlice(arr2, 1, 5).(*array.Int32)
	defer sli.Release()

	// slice = [2 3 (null) 5]
	fmt.Printf("slice = %v\n", sli)

	dtype := arrow.StructOf([]arrow.Field{
		{Name: "Name", Type: arrow.ListOf(arrow.PrimitiveTypes.Uint8)},
		{Name: "Age", Type: arrow.PrimitiveTypes.Int32},
		{Name: "Langs", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}...)

	str := array.NewStructBuilder(mem, dtype)
	defer str.Release()

	fmt.Printf("str = %#v\n", str)
}
