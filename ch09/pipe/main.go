package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)
	var w sync.WaitGroup
	w.Add(2)

	go stepOne(channel, &w)
	go stepTwo(channel, &w)

	w.Wait()
	fmt.Println("main complete")
}

// step one
//
// write message to channel
func stepOne(channel chan<- string, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Println("step 1 start")
	fmt.Println("step 1 end")
	channel <- "Complete!"
}

// step two
//
// read message from channel
func stepTwo(channel <-chan string, w *sync.WaitGroup) {
	defer w.Done()

	message := <-channel
	fmt.Printf("received message: %s\n", message)

	fmt.Println("step 2 start")
	fmt.Println("sten 2 end")
}
