package main

import "testing"

func TestHoge(t *testing.T) {
	a := 1
	b := 1
	c := a + b
	if c != 2 {
		t.Fatal("unexpected value")
	}
}
