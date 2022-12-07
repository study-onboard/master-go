package commander

import (
	"fmt"
	"log"
	"os"
)

// command handler
// if continue to run, return true, other wish return false
type commandHandler func(args []string)

// handler mapping
var handlerMapping map[string]commandHandler = make(map[string]commandHandler)

// add command handler
func AddCommandHandler(command string, handler commandHandler) {
	log.Printf("Add handler for command %q: %v\n", command, handler)
	handlerMapping[command] = handler
}

// handle command
func HandleCommand(command string, arguments []string) {
	handler, ok := handlerMapping[command]
	if !ok {
		fmt.Printf("command not found: %s\n", command)
		os.Exit(1)
	}
	handler(arguments)
}
