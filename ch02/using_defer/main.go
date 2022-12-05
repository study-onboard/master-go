package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFilePath := "./output.log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open log file failed: %v\n", err)
		os.Exit(10)
	}
	defer logFile.Close()

	logger := log.New(logFile, "Using defer - ", log.LstdFlags|log.Lshortfile)
	doSomeThings(logger)
}

func doSomeThings(logger *log.Logger) {
	logger.Println("-------------------------start---------------------------------")
	defer logger.Println("--------------------------end----------------------------------")

	messages := []string{
		"Show me the money",
		"How do you turn this on",
		"Good good study, day day up",
		"Make love, no war!",
	}
	for i, message := range messages {
		logger.Printf(" %d - %s", i, message)
	}
}
