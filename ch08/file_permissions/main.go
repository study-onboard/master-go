package main

import (
	"fmt"
	"os"
)

func main() {

	const filePath = "./main.go"

	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("get file state failed: %s\n", err.Error())
		os.Exit(1)
	}

	mode := info.Mode()
	fmt.Printf("mode is: %s\n", mode.String())
}
