package main // import "try_go_http2"

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	res, err := http.Get("https://localhost:3000")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Printf("protocol version: %s\n", res.Proto)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}
}
