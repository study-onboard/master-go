package main

import "fmt"

func main() {
	num1, num2, num3 := generate()
	fmt.Printf("num1: %d, num2: %d, num3: %d\n", num1, num2, num3)
}

func generate() (int, int, int) {
	var v1, v2, v3 int = 1, 2, 3
	v1, v2, v3 = 4, 3, 2
	return v1, v2, v3
}
