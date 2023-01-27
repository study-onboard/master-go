package pipe

import (
	"github.com/google/uuid"
)

// person
type Person struct {
	Id   string
	Name string
}

// create person
func NewPerson(name string) *Person {
	return &Person{
		Id:   uuid.NewString(),
		Name: name,
	}
}
