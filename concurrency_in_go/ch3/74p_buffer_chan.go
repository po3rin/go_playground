package ch3

import (
	"bytes"
	"fmt"
	"os"
)

func P74() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 4; i++ {
			fmt.Fprintf(&stdoutBuff, "Sdnding: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Println("dd")
		fmt.Fprintf(&stdoutBuff, "Recieve: %v.\n", integer)
	}
}
