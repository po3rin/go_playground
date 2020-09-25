package main

import "fmt"

func main() {
	m := map[int][]string{
		1: []string{"a", "b", "c", "d"},
		2: []string{"c", "b", "a", "d"},
		3: []string{"a", "b", "d", "c"},
		4: []string{"c", "a", "d", "b"},
	}
	f := map[string][]int{
		"a": []int{1, 2, 3, 4},
		"b": []int{2, 1, 4, 3},
		"c": []int{2, 3, 1, 4},
		"d": []int{1, 4, 3, 2},
	}

	matching := stableMatching(m, f)
	fmt.Println(matching)
}

// TODO: imp
func stableMatching(m map[int][]string, f map[string][]int) []string {
	matching := make([]string, len(m))
	return matching
}
