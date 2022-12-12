package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)

	if err := process(&w); err != nil {
		fmt.Printf("process error: %s\n", err.Error())
		os.Exit(1)
	}

	w.Wait()
	fmt.Println("main complete!")
}

func process(w *sync.WaitGroup) error {
	defer w.Done()

	processChannel := make(chan int)
	go func(processChannel chan<- int) {
		fmt.Println("start processing..........")
		time.Sleep(5 * time.Second)
		fmt.Println("process complete!")
		processChannel <- 0
	}(processChannel)

	select {
	case <-processChannel:
		close(processChannel)
		return nil

	case <-time.After(4 * time.Second):
		close(processChannel)
		return errors.New("process timeout")
	}
}
