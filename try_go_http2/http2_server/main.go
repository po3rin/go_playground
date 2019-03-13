package main // import "try_go_http2"

import (
	"io"
	"log"
	"net/http"
	"path/filepath"

	"try_go_http2/handler"
)

type flushWriter struct {
	w io.Writer
}

func (fw flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	// Flush - send the buffered written data to the client
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	return
}

func echoHandle(w http.ResponseWriter, r *http.Request) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	// Copy from the request body to the response writer and flush
	w.Write([]byte(r.Proto))
}

func main() {
	certFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.crt")
	keyFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.key")

	http.HandleFunc("/reqinfo", handler.ReqInfoHandler)
	// http.HandleFunc("/proto/flush", echoHandle)
	http.HandleFunc("/clock", handler.ClockStreamHandler)

	err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
