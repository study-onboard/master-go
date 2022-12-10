package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("name-%d", i)
		waitGroup.Add(1)
		go process(&waitGroup, name)
	}

	waitGroup.Wait()
	fmt.Println("All complete!")
}

// process - goroutine
func process(waitGroup *sync.WaitGroup, name string) {
	defer waitGroup.Done()

	fmt.Printf("%s processing.......\n", name)
	fmt.Printf("%s process complete!\n", name)
}
