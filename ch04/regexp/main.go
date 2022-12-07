package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "My cloud node's IP is 192.168.3.99, go to launch it and take up for accessable. and go to 192.168.4.34 to shutdown it"
	fmt.Printf("text: %s\n", text)

	pattern := regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)
	if pattern.MatchString(text) {
		fmt.Println("Matched!")
		matches := pattern.FindAllString(text, 2)
		fmt.Println(matches)
	} else {
		fmt.Println("No match!!")
	}
}
