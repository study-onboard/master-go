package main

import "fmt"

func main() {
	group()
}

func group() {
	fmt.Println("Inside group function")
	defer func() {
		// in recover
		if c := recover(); c != nil {
			fmt.Printf("Recover inside group function - %v\n", c)
		} else {
			fmt.Println("Existing group function")
		}
	}()

	fmt.Println("Calling to item function")
	item()
	fmt.Println("Item function called")

	fmt.Println("Do some things in group function")
}

func item() {
	fmt.Println("Inside item function")
	panic("Panic from item function")
	fmt.Println("Exiting item function")
}
