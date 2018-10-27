package main

import (
	"go-playground/concurrency_in_go/ch3"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// msg := "hello"
	// sayHello := func() {
	// 	defer wg.Done()
	// 	msg = "mmm"
	// }
	// wg.Add(1)
	// go sayHello()
	// wg.Wait()
	// fmt.Println(msg)

	// for _, salution := range []string{"hello", "greeting", "good bye"} {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(salution)
	// 	}()
	// }
	// wg.Wait()

	// for _, salution := range []string{"hello", "greeting", "good bye"} {
	// 	wg.Add(1)
	// 	go func(salution string) {
	// 		defer wg.Done()
	// 		fmt.Println(salution)
	// 	}(salution)
	// }
	// wg.Wait()

	// ch3.P44()

	// ch3.P53()

	// ch3.P56()

	ch3.P59()
}
