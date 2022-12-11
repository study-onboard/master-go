package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	queue := make(chan string)

	w.Add(3)
	go process("A", queue, &w)
	go process("B", queue, &w)
	go process("C", queue, &w)
	go prepare(queue)

	w.Wait()

	fmt.Println("main complete!!")
}

// prepare
func prepare(queue chan<- string) {
	fmt.Println("preparing.....")
	time.Sleep(time.Second)
	fmt.Println("prepare complete")
	queue <- "complete!!"
	close(queue)
}

// process
func process(name string, queue <-chan string, w *sync.WaitGroup) {
	defer w.Done()
	message := <-queue
	if message == "" {
		fmt.Printf("process %s received queue close signal.\n", name)
	} else {
		fmt.Printf("process %s received message: %s\n", name, message)
	}
}
