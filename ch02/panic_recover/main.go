package main

import "fmt"

func main() {
	group()
	fmt.Println("group function complete!")

	success := managerHandle()
	if success {
		fmt.Println("manager handler function call successfuly")
	} else {
		fmt.Println("manager handler function call failed")
	}
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

func managerHandle() bool {
	fmt.Println("Into manager handler function")

	defer func() {
		if c := recover(); c != nil {
			fmt.Printf("Into defer recover function block - %v\n", c)
		} else {
			fmt.Println("Existing manager handler function")
		}
	}()

	fmt.Println("Calling staff handler function")
	staffHandle()
	fmt.Println("Staff handler fucntion called")

	fmt.Println("Do some things in manager handler function")

	return true
}

func staffHandle() {
	fmt.Println("Into staff handler function")
	panic("Panic from staff handler function")
	fmt.Println("Exiting staff handler function")
}
