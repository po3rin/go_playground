package ch3

import (
	"fmt"
)

func P69() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			fmt.Printf("=======%+v\n", i)
			intStream <- i
		}
	}()
	// チャネルが閉じたらループ終了。チャネル受け取ったら回す。
	for integer := range intStream {
		fmt.Printf("%v \n", integer)
	}

}
