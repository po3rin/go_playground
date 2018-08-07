// 関数によって処理の差し替えを実現し、さらにインターフェイスも定義するケースで変更に強くする
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

// 差し替え可能 ===========

type DoStruct struct {
	X, Y int
}

func (d DoStruct) Call() {
	fmt.Println(d.X)
	fmt.Println(d.Y)
}

// ======================
