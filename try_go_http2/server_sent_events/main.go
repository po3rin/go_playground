package main // import "try_go_http2"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var html []byte

func init() {
	h, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	html = h
}

func handlerHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
}

func handlerNum(w http.ResponseWriter, r *http.Request) {
	flusher, _ := w.(http.Flusher)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	go func() {
		cnt := 1
		for {
			select {
			case <-t.C:
				fmt.Fprintf(w, "data: %d\n\n", cnt)
				cnt++
				flusher.Flush()
			default:
			}
		}
	}()

	notify := w.(http.CloseNotifier).CloseNotify()
	<-notify

	log.Println("コネクションが閉じました")
}

func main() {
	certFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.crt")
	keyFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.key")

	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/num", handlerNum)

	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", certFile, keyFile, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
