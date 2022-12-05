package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatalf("Could not open logger: %s\n", err.Error())
	} else {
		log.SetOutput(sysLogger)
	}

	log.Println("How do you turn this on?")
}
