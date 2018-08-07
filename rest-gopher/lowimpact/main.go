package main

import "fmt"

func main() {
	Do(NewHelloFunc())
}
func Do(d DoInterface) {
	d.Call()
}

type DoInterface interface {
	Call()
}

type DoFunc func()

func (d DoFunc) Call() {
	d()
}

func NewHelloFunc() DoFunc {
	f := func() {
		fmt.Println("hello")
	}
	return DoFunc(f)
}
