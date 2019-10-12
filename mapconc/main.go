package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	m := make(map[string]int)

	go func() {
		for i := 0; i < 100; i++ {
			m[strconv.Itoa(i)] = i // write
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i, m[strconv.Itoa(i)]) // read
		}
	}()

	time.Sleep(time.Second * 3)

}
