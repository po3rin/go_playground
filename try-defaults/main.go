package main

import (
	"fmt"

	"github.com/creasty/defaults"
)

type Sample struct {
	Name string `default:"John Smith"`
	Age  int    `default:"27"`
}

func main() {
	obj := &Sample{
		Name: "taro",
	}
	if err := defaults.Set(obj); err != nil {
		panic(err)
	}
	fmt.Println(obj) //&{taro 27}
}
