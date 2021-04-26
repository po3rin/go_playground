package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Println("Request error", err)
	}

	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println("Request error", err)
	}
	fmt.Println(ctx.Err())
}
