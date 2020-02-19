package main

import "fmt"

func sort(a []int) []int {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = key
	}
	return a
}

func main() {
	instance := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(instance)
	solved := sort(instance)
	fmt.Println(solved)
}
