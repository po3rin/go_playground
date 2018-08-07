package main

import "fmt"

func main() {
	Do(NewHelloFunc())
}
func Do(f DoFunc) {
	// any process ...
	f()
	// any process ...
}

type DoFunc func()

func NewHelloFunc() DoFunc {
	f := func() {
		fmt.Println("hello")
	}
	return DoFunc(f)
}
