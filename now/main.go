package main

import (
	"fmt"
	"time"
)

func main() {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now()
	fmt.Println(now)
	fmt.Println(time.Date(now.Year(), now.Month(), now.Day(), 7, 0, 0, 0, time.Local).AddDate(-1, 0, 0))
}
