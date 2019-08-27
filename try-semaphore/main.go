package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

func main() {
	ts, err := filter()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ts)
}

func filter() ([]int, error) {
	ts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ctx := context.Background()
	maxWorkers := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(maxWorkers))

	var mu sync.Mutex
	re := ts[:0]
	for _, t := range ts {

		if err := sem.Acquire(ctx, 1); err != nil {
			break
		}

		go func(t int) {
			defer sem.Release(1)

			mu.Lock()
			re = append(re, t)
			mu.Unlock()

		}(t)
	}

	// wg.Wait()
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	return re, nil
}
