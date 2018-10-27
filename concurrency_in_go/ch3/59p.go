package ch3

import (
	"fmt"
	"sync"
)

func P59() {
	var count int
	increment := func() { count++ }
	decriment := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decriment)

	fmt.Println(count)
}
