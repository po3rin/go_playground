package main // import "try_go_http2"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./img/golang.png")
	if err != nil {
		panic(err)
	}
}

func handlerHTML(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "image/png")
	w.Write(image)
}

func handlerByte(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ww")
	w.Write([]byte("Hello Again!"))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/byte", nil)
	} else {
		fmt.Println("no push")
	}
	// Send response body
	w.Write([]byte("Hello"))
}

func main() {
	certFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.crt")
	keyFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.key")

	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/image", handlerImage)
	http.HandleFunc("/byte", handlerByte)
	http.HandleFunc("/index", handlerIndex)

	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", certFile, keyFile, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
