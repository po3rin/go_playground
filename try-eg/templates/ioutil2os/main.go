package main

import (
	"io/ioutil"
	"os"
)

func before(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
func after(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
