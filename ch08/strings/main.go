package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := "show me the 钱"
	reader := strings.NewReader(text)
	// 钱 will length 3
	fmt.Printf("Length: %d\n", reader.Len())
	reader.WriteTo(os.Stdout)
	fmt.Println()
}
