package main

import (
	"log"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("waaa")
	})

	log.Println("ready")
	err := http.ListenAndServe(
		":3000", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("start")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Println("done")
}
