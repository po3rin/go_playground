package main // import "try_go_http2"

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"

	"golang.org/x/net/http2"
)

func main() {
	certFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.crt")
	keyFile, _ := filepath.Abs("/Users/hiromunakamura/.cert/localhost.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// var buf bytes.Buffer
		// e := hpack.NewEncoder(&buf)

		// e.WriteField(hpack.HeaderField{
		// 	Name:  ":status",
		// 	Value: "204",
		// })
		buf := new(bytes.Buffer)
		fr := http2.NewFramer(buf, buf)

		var streamID uint32 = 1<<24 + 2<<16 + 3<<8 + 4
		fr.WriteData(streamID, true, []byte("Hello"))
	})

	err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
