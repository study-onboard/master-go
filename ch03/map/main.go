package main

import (
	"fmt"

	"github.com/study-onboard/master-go/common/ui"
)

func main() {
	department := map[string]string{
		"Boss":    "Lana",
		"CEO":     "Qing",
		"Manager": "Kut",
	}

	ui.PrintTitle("Simply print map")
	for position, name := range department {
		fmt.Printf("%s: %s\n", position, name)
	}
	ui.PrintEnd()

	// delete key
	delete(department, "CEO")
	ui.PrintTitle("After key deleted")
	for position, name := range department {
		fmt.Printf("%s: %s\n", position, name)
	}
	ui.PrintEnd()

	// check key exists
	ui.PrintTitle("Check `CEO` key")
	if _, exists := department["CEO"]; exists {
		fmt.Println("key 'CEO' exists")
	} else {
		fmt.Println("key 'CEO' not found")
	}
	ui.PrintEnd()
}
