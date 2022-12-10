package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)

	var prepareWaitGroup sync.WaitGroup
	prepareWaitGroup.Add(3)

	var startWaitGroup sync.WaitGroup
	startWaitGroup.Add(3)

	go process("A", channel, &prepareWaitGroup, &startWaitGroup)
	go process("B", channel, &prepareWaitGroup, &startWaitGroup)
	go process("C", channel, &prepareWaitGroup, &startWaitGroup)

	prepareWaitGroup.Wait()
	close(channel)

	startWaitGroup.Wait()
	fmt.Println("main, bye bye!")
}

func process(
	name string,
	channel <-chan int,
	prepareWaitGroup *sync.WaitGroup,
	startWaitGroup *sync.WaitGroup,
) {
	defer startWaitGroup.Done()

	fmt.Printf("%s waiting signal\n", name)
	prepareWaitGroup.Done()

	<-channel
	fmt.Printf("%s touched\n", name)
}
