package main

import (
	"log"
	"os"
)

const LOGFILE string = "/tmp/ch01_custom_log_file.log"

func main() {
	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// create custom logger
	iLog := log.New(logFile, "Custom Log - ", log.LstdFlags)

	// print logs by custom logger
	iLog.Println("Show me the money!")
	iLog.Println("How do you turn this on?")

	// set log output the source file and source code line number
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Good good study, day day up!!")
}
