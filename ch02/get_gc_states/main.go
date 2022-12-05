package main

import (
	"fmt"
	"runtime"
)

func main() {
	printMemStats()

	for i := 0; i < 10; i++ {
		buf := make([]byte, 5000000000)
		if buf == nil {
			fmt.Println("Can not alloc bytes.")
		}
	}
	printMemStats()
}

func printMemStats() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Printf("alloc: %v\n", stats.Alloc)
	fmt.Printf("total alloc: %v\n", stats.TotalAlloc)
	fmt.Printf("heap alloc: %v\n", stats.HeapAlloc)
	fmt.Printf("num gc: %v\n", stats.NumGC)
	fmt.Println("---------------------------------------")
}
