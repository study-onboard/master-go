package main

import "fmt"

func main() {
	fmt.Println("show me the", "money")
	fmt.Printf("show me the %s\n", "money")
	fmt.Print("show me the money\n")
	fmt.Println(fmt.Sprintf("show me the %s", "money"))
}
