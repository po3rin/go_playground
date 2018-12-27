package main // import "try-tag"

import (
	"encoding/json"
	"fmt"

	"github.com/creasty/defaults"
)

type Str struct {
	Name string
}

type Sample struct {
	SliceByJSON []int          `default:"[]"` // Supports JSON format
	MapByJSON   map[string]int `default:"{}"`
	Str         *Str           `default:"{}"`
}

func main() {
	obj := &Sample{}
	if err := defaults.Set(obj); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", obj)

	bytes, _ := json.Marshal(obj)
	// want [] ...
	fmt.Println(string(bytes))
}
