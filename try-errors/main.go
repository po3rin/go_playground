package main

import (
	"errors"
	"fmt"
)

func f() error {
	return errors.New("new errors")
}

func g() error {
	err := f()
	if err != nil {
		return fmt.Errorf("failed to call f: %w", err)
	}
	return nil
}

func main() {
	err := g()
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
