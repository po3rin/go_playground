package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"cuelang.org/go/cue"
)

func main() {
	body, err := ioutil.ReadFile("people.cue")
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		os.Exit(1)
	}

	var r cue.Runtime
	ins, err := r.Compile("people", body)
	if err != nil {
		fmt.Printf("compile test instance: %v\n", err)
		os.Exit(1)
	}

	people := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "a1",
			Age:  23,
		},
		{
			Name: "a2",
			Age:  24,
		},
		{
			Name: "a3",
			Age:  25,
		},
	}

	ins, err = ins.Fill(people, "people")
	if err != nil {
		fmt.Printf("fill: %v\n", err)
		os.Exit(1)
	}

	json, err := ins.Value().MarshalJSON()
	if err != nil {
		fmt.Printf("marshal json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(json))
}
