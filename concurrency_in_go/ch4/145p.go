package ch4

import "fmt"

func P145() {
	type foo int
	type bar int

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2

	fmt.Println(m)
}
