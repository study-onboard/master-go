package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup

	tasks := []string{
		"A", "B", "C",
	}
	c := make(chan int)
	for _, task := range tasks {
		w.Add(1)
		go runTask(task, c, &w)
	}

	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		go func(number int) {
			c <- number
		}(i)
	}

	w.Wait()
	fmt.Println("main complete!")
}

func runTask(name string, c <-chan int, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Printf("task %s waiting number............\n", name)
	number := <-c
	fmt.Printf("task %s received number: %d\n", name, number)

	for i := 0; i < 999; i++ {
		fmt.Printf("task %s complete: %d\n", name, number)
	}
}
