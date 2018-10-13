package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	return
}
