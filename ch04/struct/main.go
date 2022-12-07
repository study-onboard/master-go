package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	initializeCase()
	structPointer()
}

type person struct {
	id       string
	name     string
	position string
	age      uint
	comment  *string
}

// covert person to string
func (p *person) String() string {
	convertStringPoint := func(value *string) string {
		if value == nil {
			return "<nil>"
		} else {
			return *value
		}
	}
	return fmt.Sprintf(
		"id: %s, name: %s, position: %s, age: %d, comment: %s",
		p.id,
		p.name,
		p.position,
		p.age,
		convertStringPoint(p.comment),
	)
}

// convert value to string pointer
func toStringPointer(value string) *string {
	return &value
}

// initialize case
func initializeCase() {
	// style 1
	p := person{
		uuid.NewString(),
		"kut",
		"manager",
		44,
		nil,
	}
	fmt.Printf("person: %s\n", p.String())

	// style 2
	p = person{
		id:       uuid.NewString(),
		name:     "lana",
		position: "boss",
		age:      42,
		comment:  toStringPointer("My God"),
	}
	fmt.Printf("person: %s\n", p.String())
}

func structPointer() {
	// change person field
	changePersonComment := func(pt *person, comment string) {
		pt.comment = &comment
	}

	// build person by return pt
	buildPerson := func(name, position, comment string, age uint) *person {
		return &person{
			uuid.NewString(),
			name,
			position,
			age,
			&comment,
		}
	}

	// build person by new method
	buildPerson2 := func(name, position string, age uint, comment string) *person {
		result := new(person)
		result.id = uuid.NewString()
		result.name = name
		result.position = position
		result.age = age
		result.comment = &comment
		return result
	}

	pt := buildPerson("kut", "manager", "IT Manager", 44)
	fmt.Printf("pt: %s\n", pt.String())

	changePersonComment(pt, "Real name is Kut Zhang")
	fmt.Printf("after change name, pt: %s\n", pt.String())

	pt = buildPerson2("Lana", "Boss", 42, "Good bleed")
	fmt.Printf("pt2: %s\n", pt.String())
}
