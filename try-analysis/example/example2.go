package example

import "fmt"

type s struct{}

func (s s) Println(x string) {}

func Print() {
	fmt.Println("xxx")

	{
		fmt := s{}
		fmt.Println("yyy")
	}

	fmt.Println("xxx")
}
