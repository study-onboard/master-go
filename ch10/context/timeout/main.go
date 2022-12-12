package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := []string{
		"A", "B",
	}
	var w sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)

	w.Add(len(tasks))
	for _, task := range tasks {
		go runTask(task, &w, ctx)
	}

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	w.Wait()
	fmt.Println("main complete!")
}

func runTask(name string, w *sync.WaitGroup, ctx context.Context) {
	defer w.Done()

	fmt.Printf("task %s started\n", name)

	select {
	case <-ctx.Done():
		fmt.Printf("task %s failed: %s\n", name, ctx.Err())

	case <-time.After(5 * time.Second):
		fmt.Printf("task %s complete!\n", name)
		return
	}
}
