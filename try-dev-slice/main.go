package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func at(s reflect.SliceHeader, i int) unsafe.Pointer {
	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func main() {
	a := [...]int{10, 20, 30}
	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&a[0])), Len: 2, Cap: 3}
	*(*int)(at(s, 1)) = 100 // unsafe.Pointerを*intに変換して代入している
	fmt.Println(a)
}
