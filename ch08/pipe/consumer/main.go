package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/study-onboard/master-go/ch08/pipe"
)

// read persons JSON from stdin
func main() {
	decoder := json.NewDecoder(os.Stdin)
	var persons []pipe.Person
	decoder.Decode(&persons)

	fmt.Printf("received %d persons.\n", len(persons))
	for i, person := range persons {
		fmt.Printf("%d. %s - %s\n", i, person.Id, person.Name)
	}
}
