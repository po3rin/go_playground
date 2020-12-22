package main

import (
	"fmt"
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Bookmark struct {
	Name string
}

var bs = []Bookmark{
	{"Twitter"},
	{"Netflex"},
	{"AWS"},
}

func main() {
	id, err := fuzzyfinder.Find(
		bs,
		func(i int) string {
			return bs[i].Name
		},
		fuzzyfinder.WithPreviewWindow(
			func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf("Name: %s", bs[i].Name)
			},
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", bs[id])
}
