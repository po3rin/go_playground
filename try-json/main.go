package main // import "try-json"

import "fmt"

type User struct {
	Name Name
	Age  int
}

type Name struct {
	First string
	Last  string
}

func main() {
	fmt.Println("vim-go")
}
