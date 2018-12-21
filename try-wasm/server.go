package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("a", ":5555", "address:port")
	flag.Parse()
	log.Printf("listening on %q\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir("."))))
}
