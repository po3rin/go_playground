package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"cuelang.org/go/cue"
)

func main() {
	b, err := ioutil.ReadFile("./query.cue")
	if err != nil {
		log.Fatal(err)
	}

	config := string(b)
	var r cue.Runtime

	ins, err := r.Compile("test", config)
	if err != nil {
		log.Fatal(err)
	}

	ins, err = ins.Fill([]string{"バナナ", "アレルギー"}, "queries")
	if err != nil {
		log.Fatal(err)
	}
	ins, err = ins.Fill(true, "useDataSort")
	if err != nil {
		log.Fatal(err)
	}
	ins, err = ins.Fill(3, "size")
	if err != nil {
		log.Fatal(err)
	}

	json, err := ins.Value().MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
