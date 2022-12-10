package main

import (
	"bytes"
	"fmt"
)

func main() {
	var data []byte
	buffer := bytes.NewBuffer(data)
	buffer.Write([]byte("show me the money\n"))
	buffer.Write([]byte("how do you turn this on\n"))
	buffer.Write([]byte("good good study, day day up.\n"))

	result := string(buffer.Bytes())
	fmt.Println("result: ")
	fmt.Println("--------------------------------------------")
	fmt.Printf(result)
}
