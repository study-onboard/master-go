package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var w sync.WaitGroup
	w.Add(2)

	go taskA(ctx, &w)
	go taskB(ctx, &w)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("process goroutine canceled.")
		cancel()
	}()

	w.Wait()
	fmt.Println("main complete")
}

func taskA(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Println("task A started.")
	select {
	case <-ctx.Done():
		fmt.Printf("cancel task A: %s\n", ctx.Err())

	case <-time.After(4 * time.Second):
		fmt.Println("task A complete!")
	}
}

func taskB(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()

	fmt.Println("task B started.")
	select {
	case <-ctx.Done():
		fmt.Printf("cancel task B: %s\n", ctx.Err())

	case <-time.After(4 * time.Second):
		fmt.Println("task B complete!")
	}
}
