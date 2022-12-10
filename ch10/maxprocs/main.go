package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("max procs: %d\n", runtime.GOMAXPROCS(0))
}
