package ch4

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func P139() {

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // <1>
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreetingctx(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel() // <2>
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewellctx(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreetingctx(ctx context.Context) error {
	greeting, err := genGreetingctx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewellctx(ctx context.Context) error {
	farewell, err := genFarewellctx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreetingctx(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // <3>
	defer cancel()

	switch locale, err := localectx(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewellctx(ctx context.Context) (string, error) {
	switch locale, err := localectx(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func localectx(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err() // <4>
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
