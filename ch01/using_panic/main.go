package main

import (
	"errors"
	"os"
)

func main() {
	err := errors.New("Some error by codes!!")
	panic(err)
	os.Exit(10)
}
