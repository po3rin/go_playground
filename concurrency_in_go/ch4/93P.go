package ch4

import "fmt"

func P93() {
	// gorutine leak.
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("do work exited")
			defer close(completed)
			for i := range strings {
				fmt.Println(i)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("done")
}
