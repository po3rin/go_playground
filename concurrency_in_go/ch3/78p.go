package ch3

import (
	"fmt"
	"time"
)

// P78 channelの所有権について
func P78() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			defer fmt.Println("wwww")
			for i := 0; i <= 5; i++ {
				resultStream <- i
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}()
		fmt.Println("aaaa")
		return resultStream
	}

	resultStream := chanOwner()
	fmt.Println("bbbb")
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done Received")
}

// // output
// aaaa
// bbbb
// Received: 0
// 0
// 1
// Received: 1
// 2
// Received: 2
// 3
// Received: 3
// 4
// Received: 4
// 5
// Received: 5
// wwww
// Done Received
