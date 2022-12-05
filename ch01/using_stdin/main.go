package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := os.Stdin
	defer stdin.Close()

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "quit" || input == "exit" {
			break
		}

		fmt.Println("-- ", input)
	}

	fmt.Println("-- bye bye!!")
}
