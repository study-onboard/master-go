package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var w sync.WaitGroup
var resource string

func main() {
	// init resource
	resource = "main"
	fmt.Printf("resource in main: %s\n", resource)

	// define tasks
	tasks := []string{
		"A", "B", "C", "D",
	}

	// run tasks
	w.Add(len(tasks))
	for _, t := range tasks {
		go task(t)
	}

	// wait for all goroutine
	w.Wait()
	fmt.Println("main complete!")
}

// task
func task(name string) {
	defer w.Done()

	mutex.Lock()
	defer mutex.Unlock()

	prev := resource
	resource = name
	fmt.Printf("task %s - prev: %s - current: %s\n", name, prev, resource)
}
