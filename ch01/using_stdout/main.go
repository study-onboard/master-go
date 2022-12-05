package main

import (
	"io"
	"os"
)

func main() {
	var message string

	if len(os.Args) > 1 {
		message = os.Args[1]
	} else {
		message = "Please enter the message!"
	}

	io.WriteString(os.Stdout, message)
	io.WriteString(os.Stdout, "\n")

	io.WriteString(os.Stderr, "Good bye!")
	io.WriteString(os.Stderr, "\n")
}
