package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	signalChannel := make(chan int)

	var w sync.WaitGroup
	w.Add(3)

	go startTask("A", signalChannel, &w)
	go startTask("B", signalChannel, &w)
	go startTask("C", signalChannel, &w)

	fmt.Println("start process main preparing....")
	time.Sleep(3 * time.Second)
	fmt.Println("main prepare complete")
	close(signalChannel)

	w.Wait()
	fmt.Println("main complete")
}

func startTask(name string, signalChannel <-chan int, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Printf("task %s waiting signal to run........\n", name)
	<-signalChannel

	fmt.Printf("task %s running..............\n", name)
	time.Sleep(time.Second)
	fmt.Printf("task %s compete!\n", name)
}
