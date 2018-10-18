package main

import (
	"fmt"
	"time"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	println("Err")
	// }
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	println(scanner.Text())
	// }

	for i := 0; ; {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		// i++
	}

	// for {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("ww")
	// }
}
