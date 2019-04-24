package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/creasty/defaults"
)

type Sample struct {
	Name     string   `default:"John Smith"`
	Age      int      `default:"27"`
	Contents []string `default:"[]"`
}

func main() {
	obj := &Sample{
		Name: "taro",
	}

	var buf bytes.Buffer
	b, _ := json.Marshal(obj)
	buf.Write(b)
	fmt.Println(buf.String())

	print(obj)
}

func print(a interface{}) {
	if err := defaults.Set(a); err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	b, _ := json.Marshal(a)
	buf.Write(b)
	fmt.Println(buf.String())
}
