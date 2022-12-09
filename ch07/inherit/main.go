package main

import "fmt"

func main() {
	s := &staff{
		human: human{
			name: "Kut Zhang",
		},
		company: "Sanlea",
	}
	s.walk()
	s.work()
}

type human struct {
	name string
}

func (h *human) walk() {
	fmt.Printf("human %q walking....\n", h.name)
}

type staff struct {
	human
	company string
}

func (s *staff) work() {
	fmt.Printf("%q staff - %q working.......\n", s.company, s.name)
}
