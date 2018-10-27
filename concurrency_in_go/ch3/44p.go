package ch3

import (
	"fmt"
	"runtime"
	"sync"
)

func P44() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGorutines = 1e4
	wg.Add(numGorutines)
	before := memConsumed()
	for i := numGorutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGorutines/1000)
}
