package main

import "fmt"

func main() {
	h := &Hello{}
	Do(h)
}

func Do(f DoFunction) {
	// any process
	f.Exec()
	// any process
}

type DoFunction interface {
	Exec()
}

type Hello struct{}

func (h *Hello) Exec() {
	fmt.Println("hello")
}
