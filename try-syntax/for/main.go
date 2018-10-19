package main

import (
	"fmt"
	"reflect"
)

func main() {
	// for loop := 0; ; {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(loop)
	// }
	a := [3]int{2, 3, 4}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a[:]))
}
