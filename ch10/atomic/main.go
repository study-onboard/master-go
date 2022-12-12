package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var resource int64

func main() {
	resource = 0
	fmt.Printf("in main, resource: %d\n", resource)

	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			w.Add(1)
			go add(&w)

			w.Add(1)
			go sub(&w)
		}
	}

	result := atomic.LoadInt64(&resource)
	fmt.Printf("result: %d\n", result)
}

func add(w *sync.WaitGroup) {
	defer w.Done()
	atomic.AddInt64(&resource, 1)
}

func sub(w *sync.WaitGroup) {
	defer w.Done()
	atomic.AddInt64(&resource, -1)
}
