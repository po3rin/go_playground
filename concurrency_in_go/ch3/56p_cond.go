package ch3

import (
	"fmt"
	"sync"
)

func P56() {

	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickedRegistered sync.WaitGroup
	clickedRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickedRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaing annoying dialog box!")
		clickedRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clickded.")
		clickedRegistered.Done()
	})

	button.Clicked.Broadcast()

	clickedRegistered.Wait()
}
