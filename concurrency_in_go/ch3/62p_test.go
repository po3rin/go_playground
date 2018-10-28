package ch3

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	deamonStarted := startNetworkDaemon2()
	// deamonStarted := startNetworkDaemon()
	deamonStarted.Wait()
}

func BenchmarkNetwokRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read : %v", err)
		}
		conn.Close()
	}
}
