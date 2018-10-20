package main

import (
	"fmt"
	"time"
)

func main() {
	for loop := 0; ; {
		time.Sleep(1 * time.Second)
		fmt.Println(loop)
	}
}
