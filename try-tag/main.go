package main // import "try-tag"

import (
	"encoding/json"
	"fmt"

	"github.com/creasty/defaults"
	"github.com/pkg/errors"
)

type Str struct {
	Name string
	Arg  []int          `default:"[]"`
	Map  map[string]int `default:"{}"`
}
type Sample struct {
	SliceByJSON []int          `default:"[]"` // Supports JSON format
	MapByJSON   map[string]int `default:"{}"`
	Str         *Str           `default:"{}"`
	Responses   interface{}    `default:"[]"`
}

func main() {
	obj := &Sample{}

	bytes, _ := json.Marshal(obj)
	fmt.Println(string(bytes))

	i := SetDefault(obj)
	bytes, _ = json.Marshal(i)
	fmt.Println(string(bytes))
}

func SetDefault(data interface{}) interface{} {
	if err := defaults.Set(data); err != nil {
		fmt.Println(errors.Wrap(err, "failed to set default"))
	}
	return data
}
