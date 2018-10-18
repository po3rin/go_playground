package main

import (
	"fmt"
	"time"
)

// Looper - for origin loop func.
type Looper struct {
	val int
}

// NewListLooper -
func NewListLooper() *Looper {
	return &Looper{}
}

// Loop - loop slice.
func (l *Looper) Loop() bool {
	for i := 1; ; {
		// fmt.Println(i)
		time.Sleep(1 * time.Second)
		l.val = i
		i++
		fmt.Println(i)
		return true
	}
}

func main() {
	l := NewListLooper()
	for l.Loop() {
		fmt.Println(l.val)
	}
}
