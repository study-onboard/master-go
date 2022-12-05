package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	printEnv()

	if checkGoVersion() == false {
		log.Panicf("Need Go version 1.8 or higher!")
	}

	fmt.Println("App go on!!")
}

func printEnv() {
	log.Printf("You are using %s on a %s machine", runtime.Compiler, runtime.GOARCH)
	log.Printf("Using Go version: %s", runtime.Version())
	log.Printf("Number of CPUs: %d", runtime.NumCPU())
	log.Printf("Number of goroutines: %d", runtime.NumGoroutine())
}

func checkGoVersion() bool {
	version := runtime.Version()
	parts := strings.Split(version, ".")
	major := string(parts[0][2])
	minor := parts[1]

	majorNumber, _ := strconv.ParseInt(major, 10, 32)
	minorNumber, _ := strconv.ParseInt(minor, 10, 32)

	if majorNumber < 1 {
		return false
	}

	if majorNumber < 1 || (majorNumber == 1 && minorNumber < 8) {
		return false
	}

	return true
}
