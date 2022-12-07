package handlers

import (
	"fmt"

	"github.com/study-onboard/master-go/ch04/commander/v1"
)

func init() {
	commander.AddCommandHandler("hello", func(arguments []string) {
		fmt.Println("Hello World!")
	})
}
