package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	operateChannel := make(chan string)
	controlChannel := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go dispatcher(operateChannel, controlChannel, &waitGroup)

	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Read message failed: %s\n", err.Error())
			controlChannel <- 0
		}
		message = message[:len(message)-1]
		if message == "" {
			continue
		}

		fmt.Println("---------------------------------------------------")
		fmt.Printf("Received message: %s\n", message)

		if message == "quit" {
			controlChannel <- 0
			break
		}

		operateChannel <- message
	}

	fmt.Println("main complete!")
}

// dispatcher
func dispatcher(
	operateChannel <-chan string,
	controlChannel <-chan int,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	for {
		select {
		case message := <-operateChannel:
			go handleMessage(message)

		case <-controlChannel:
			fmt.Println("dispatcher complete!")
			return

		case currentTime := <-time.After(4 * time.Second):
			fmt.Printf("current time: %v\n", currentTime)
		}

	}
}

// handle message
func handleMessage(message string) {
	fmt.Printf("Handling message: %s\n\n", message)
}
