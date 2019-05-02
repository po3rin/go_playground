package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("go.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, src)

	file, err := os.Create("out.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}
