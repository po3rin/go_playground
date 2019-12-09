package main

import (
	"go-playground/concurrency_in_go/ch3"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println(runtime.NumCPU())
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
	// ch3.P59()
	// ch3.P60()
	// ch3.P69()
	ch3.P74()
	// ch3.P78()

	// ch4.P89()
	// ch4.P90()
	// ch4.P94()
	// ch4.P95()
	// ch4.P96()
	// ch4.P100()
	// ch4.P101()
	// ch4.P106()
	// ch4.P119()
	// ch4.P123()
	// ch4.P139()
	// ch4.P145()

	// ch5.P153()
	// ch5.P165()
	// ch5.P167()
	// ch5.P183()
}
