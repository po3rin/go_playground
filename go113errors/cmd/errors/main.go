package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = &SampleError{
	statusCode: 404,
	level:      "Error",
	msg:        "not found",
}

type SampleError struct {
	level      string
	statusCode int
	msg        string
}

func (e *SampleError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.statusCode, e.msg)
}

func main() {
	err := func1()
	if err != nil {
		var sampleErr *SampleError
		if errors.As(err, &sampleErr) {
			switch sampleErr.level {
			case "Fatal":
				fmt.Printf("Fatal！ %v\n", sampleErr)
			case "Error":
				fmt.Printf("Error！ %v\n", sampleErr)
			case "Warning":
				fmt.Printf("Warning！ %v\n", sampleErr)
			}
		}

		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("エラーなし\n")
}

func func1() error {
	err := func2()
	if err != nil {
		return fmt.Errorf("func1 error: %w", err)
	}
	return nil
}

func func2() error {
	err := func3()
	if err != nil {
		return fmt.Errorf("func2 error: %w", err)
	}
	return nil
}
func func3() error {
	return ErrNotFound
}
