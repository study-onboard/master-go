package main

import (
	"fmt"
	"sync"
	"time"
)

var resource string
var m sync.RWMutex
var w sync.WaitGroup

func main() {
	resource = "main"
	fmt.Printf("main resource: %s\n", resource)

	readTasks := []string{
		"A", "B", "C", "D",
	}

	writeTasks := []string{
		"a", "b",
	}

	w.Add(len(readTasks) + len(writeTasks))
	for _, task := range readTasks {
		go readTask(task)
	}
	for _, task := range writeTasks {
		go writeTask(task)
	}

	w.Wait()

	fmt.Println("main complete")
}

func readTask(name string) {
	defer w.Done()

	for {
		m.RLock()
		fmt.Printf("read task %q read resource: %q\n", name, resource)
		time.Sleep(time.Second)
		m.RUnlock()
	}
}

func writeTask(name string) {
	defer w.Done()

	for {
		m.Lock()
		prev := resource
		resource = name
		fmt.Printf("write task %q - prev: %q, current: %q\n", name, prev, resource)
		m.Unlock()
	}
}
