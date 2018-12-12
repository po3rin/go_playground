package main // import "create-slice"

import (
	"fmt"
	"reflect"
	"unsafe"
)

func at(s reflect.SliceHeader, i int) unsafe.Pointer {
	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func myAppend(s reflect.SliceHeader, vs ...int) reflect.SliceHeader {
	// 新しい要素の追加
	for i := 0; i < len(vs); i++ {
		*((*int)(at(s, s.Len+i))) = vs[i]
	}
	return reflect.SliceHeader{Data: s.Data, Len: s.Len + len(vs), Cap: s.Cap}
}

func main() {
	ns := []int{10, 20, 30}
	// nsをunsafe.Pointerに変換する
	ptr := unsafe.Pointer(&ns)
	// ptrを*reflect.SliceHeaderにキャストして、それが指す値をsにいれる
	s := *(*reflect.SliceHeader)(ptr)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", s)

	b := [...]int{10, 20, 30}
	ss := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&b[0])), Len: 2, Cap: len(b)}
	*(*int)(at(ss, 0)) = 100 // unsafe.Pointerを*intに変換して代入している
	fmt.Println(b)

	c := [...]int{10, 20, 30}
	// s := a[0:2] -> [10 20]
	sss := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&c[0])), Len: 2, Cap: len(c)}
	sss = myAppend(sss, 400)
	var nss []int
	*(*reflect.SliceHeader)(unsafe.Pointer(&nss)) = sss
	fmt.Println(nss)
}
