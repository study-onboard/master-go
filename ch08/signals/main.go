package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGALRM)

	// wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	// signal listener
	go func() {
		for {
			sig := <-signalChannel
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught: ", sig)
				waitGroup.Done()
				break

			case syscall.SIGALRM:
				handleSignal(sig)
			}
		}
	}()

	// pending and wait group done
	waitGroup.Wait()
}

// handle signal
func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() caught: ", signal)
}
