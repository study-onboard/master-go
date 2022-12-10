package pipe

import (
	"github.com/google/uuid"
)

type Person struct {
	Id   string
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Id:   uuid.NewString(),
		Name: name,
	}
}
