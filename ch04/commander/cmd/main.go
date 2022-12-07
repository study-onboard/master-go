package main

import (
	"fmt"
	"os"

	"github.com/study-onboard/master-go/ch04/commander/v1"
	_ "github.com/study-onboard/master-go/ch04/commander/v1/handlers"
)

// commander command arg1 .... argN
func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Printf("usage: %s <command> <arg1...argN>", arguments[0])
		os.Exit(1)
	}

	command := arguments[1]
	commandArguments := arguments[2:]
	commander.HandleCommand(command, commandArguments)
}
