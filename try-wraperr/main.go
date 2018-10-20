package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := Bar()
	err = errors.Wrap(err, "oh noes")
	fmt.Printf("%+v\n", err)

	// output ----------------------
	// whoops
	// main.Foo
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:25
	// main.Bar
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:21
	// main.main
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:10
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:201
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1333
	// bar
	// main.Bar
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:21
	// main.main
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:10
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:201
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1333
	// oh noes
	// main.main
	// 	/Users/nakamura/hi-dev/go-playground/src/go-playground/try-wraperr/main.go:11
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:201
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1333

	errCause := errors.Cause(err)
	fmt.Println(errCause)
	// output ----------------------
	// whoops
}

func Bar() error {
	return errors.Wrap(Foo(), "bar")
}

func Foo() error {
	return errors.New("whoops")
}
