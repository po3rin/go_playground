package main

import (
	"fmt"
)

var inf = 16777619

func merge(A []int, p, q, r int) []int {
	L := make([]int, q-p+1)
	R := make([]int, r-q)

	copy(L, A[p:q+1])
	copy(R, A[q+1:r+1])

	L = append(L, inf)
	R = append(R, inf)

	var i, j int
	for k := p; k < r+1; k++ {
		if L[i] >= R[j] {
			A[k] = R[j]
			j++
			continue
		}
		A[k] = L[i]
		i++
	}

	return A
}

func sort(A []int, p, r int) []int {
	if p >= r {
		return A
	}

	q := (p + r) / 2
	A = sort(A, p, q)
	A = sort(A, q+1, r)
	A = merge(A, p, q, r)

	return A
}

func main() {
	fmt.Println("----------merge-----------")
	instance := []int{3, 4, 6, 8, 1, 2, 5, 7}
	solved := merge(instance, 0, 3, 7)
	fmt.Println(solved)

	fmt.Println("-----------merge----------")
	instance = []int{3, 4, 6, 7, 1, 2, 5}
	solved = merge(instance, 1, 3, 6)
	fmt.Println(solved)

	fmt.Println("----------sort-----------")
	instance = []int{2, 1, 7, 8, 4, 3, 5, 6}
	solved = sort(instance, 0, 7)
	fmt.Println(solved)
}
