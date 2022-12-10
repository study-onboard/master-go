package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	sendMessage("task A", "show me the money", 4, &waitGroup)
	sendMessage("task B", "how do you turn this on", 2, &waitGroup)

	waitGroup.Wait()
	fmt.Println("main complete!!!")
}

func sendMessage(name, message string, duration int, waitGroup *sync.WaitGroup) {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		c <- message
	}()
	go handler(name, c, waitGroup)
}

func handler(name string, c <-chan string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	select {
	case message := <-c:
		fmt.Printf("task %s handle message: %s\n", name, message)

		// wait channel message, 3 seconds timeout
	case <-time.After(3 * time.Second):
		fmt.Printf("task %s timeout!\n", name)
		return
	}
}
