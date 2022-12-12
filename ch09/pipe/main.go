package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)
	var w sync.WaitGroup
	w.Add(2)

	go step1(channel, &w)
	go step2(channel, &w)

	w.Wait()
	fmt.Println("main complete")
}

func step1(channel chan<- string, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Println("step 1 start")
	fmt.Println("step 1 end")
	channel <- "Complete!"
}

func step2(channel <-chan string, w *sync.WaitGroup) {
	defer w.Done()

	message := <-channel
	fmt.Printf("received message: %s\n", message)

	fmt.Println("step 2 start")
	fmt.Println("sten 2 end")
}
