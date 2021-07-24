package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
)

type Request struct {
	ID   string
	Tags []Tag
}

type Tag struct {
	ID    string
	Comma string
}

func main() {
	b, err := ioutil.ReadFile("./mlt_gauss.tmp")
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New("").Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}

	m := Request{
		ID: "aaa",
		Tags: []Tag{
			{
				ID:    "bbb",
				Comma: ",",
			},
			{
				ID: "ccc",
			},
		},
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
