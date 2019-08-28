package main

import (
	"fmt"

	unroll "github.com/po3rin/unrolling"
)

func makeBenchData(n int) []uint64 {
	ids := make([]uint64, n)
	for i := 0; i < n; i++ {
		ids[i] = uint64(i)
	}
	return ids
}

func main() {
	ids := makeBenchData(10000)
	fmt.Println(unroll.ContainsUint64(ids, 5000))
	fmt.Println(unroll.ContainsUint64Unroll2(ids, 5000))
	fmt.Println(unroll.ContainsUint64Unroll2(ids, 5000))
	fmt.Println(unroll.ContainsUint64Unroll2(ids, 5000))
	fmt.Println(unroll.ContainsUint64Unroll8WithBoundsCheck(ids, 5000))
}
