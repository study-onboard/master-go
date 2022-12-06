package main

import "fmt"

func main() {
	number := 10
	pNumber := &number
	*pNumber += 2
	fmt.Printf("number: %d\n", number)

	changeNumber(&number)
	fmt.Printf("number: %d\n", number)
}

func changeNumber(pNumber *int) {
	*pNumber += 4
}
