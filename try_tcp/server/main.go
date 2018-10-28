package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listen, _ := net.Listen("tcp", "localhost:9999")
	fmt.Println("サーバ起動@http://localhost:9999")
	for {
		conn, _ := listen.Accept()
		go func() {
			fmt.Printf("リモートアドレスは：%v", conn.RemoteAddr())

			request, _ := http.ReadRequest(bufio.NewReader(conn))
			dump, _ := httputil.DumpRequest(request, true)
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World.\n")),
			}

			response.Write(conn)
			conn.Close()

		}()
	}
}
