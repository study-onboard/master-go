package main

import (
	"fmt"
	"time"

	// import external module
	"github.com/google/uuid"
)

func main() {
	times := 0
	for {
		if times > 20 {
			break
		}
		times += 1

		fmt.Printf("UUID: %v\n", uuid.NewString())

		time.Sleep(time.Second)
	}

	fmt.Println("Good bye.")
}
