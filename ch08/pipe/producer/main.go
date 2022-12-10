package main

import (
	"encoding/json"
	"os"

	"github.com/study-onboard/master-go/ch08/pipe"
)

func main() {
	persons := []pipe.Person{
		*pipe.NewPerson("Kut Zhang"),
		*pipe.NewPerson("Lana Chang"),
		*pipe.NewPerson("Boss Liu"),
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(&persons)
}
