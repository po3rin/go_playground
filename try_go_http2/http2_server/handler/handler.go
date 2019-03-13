package handler // import "try_go_http2"

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ClockStreamHandler send time chunk.
func ClockStreamHandler(w http.ResponseWriter, r *http.Request) {
	clientGone := w.(http.CloseNotifier).CloseNotify()
	w.Header().Set("Content-Type", "text/plain")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// fmt.Fprintf(w, "# ~1KB of junk to force browsers to start rendering immediately: \n")
	// io.WriteString(w, strings.Repeat("# xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n", 13))

	for {
		fmt.Fprintf(w, "%v\n", time.Now())
		w.(http.Flusher).Flush()
		select {
		case <-ticker.C:
		case <-clientGone:
			log.Printf("Client %v disconnected from the clock", r.RemoteAddr)
			return
		}
	}
}

// ReqInfoHandler return reqest info.
func ReqInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	// fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	// fmt.Fprintf(w, "Host: %s\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	// fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	// fmt.Fprintf(w, "URL: %#v\n", r.URL)
	// fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	// fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	// fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	// fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)
}
