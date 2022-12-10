package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)
	var waitGroup sync.WaitGroup

	fmt.Println("main starting.................")

	waitGroup.Add(1)
	go produce(channel, &waitGroup)

	waitGroup.Add(1)
	go consume(channel, &waitGroup)

	waitGroup.Wait()
	fmt.Println("main complete!")
}

// produce
func produce(channel chan<- string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Writing message by producer................")
	channel <- "show me the money"
	fmt.Println("Message wrote by producer!")
	close(channel)
}

// consume
func consume(channel <-chan string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Println("Waiting message by consumer..........")
	for message := range channel {
		fmt.Printf("Received message by consumer: %s\n", message)
	}
	fmt.Println("Consumer process complete!")
}
