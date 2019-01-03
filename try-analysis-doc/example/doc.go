// Package example is example pacakge.
// Write your go code in the editor on the left and watch it previewed here on the right.
//
// Features
//
// * Supports all the GoDoc syntax
//
// * That's because this is using the actual godoc renderer compiled to WebAssembly and running in your browser!
//
// * You don't even have to give a full working sample: unresolved symbols are automagically fixed so event just a small snippet will work fine.
package example

import "fmt"

// A print "A".
func A() {
	fmt.Println("A")
}

// B struct.
type B struct {
	Name string
}
