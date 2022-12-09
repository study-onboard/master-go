package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	name := flag.String("name", "<nil>", "name flag")
	age := flag.Int("age", 20, "age flag")

	var nameFlags NamesFlag
	flag.Var(&nameFlags, "persons", "Comma-separated list")
	flag.Parse()

	fmt.Printf("%s - %d\n", *name, *age)
	fmt.Printf("persons: %v\n", nameFlags.GetNames())
}

type NamesFlag struct {
	names []string
}

func (n *NamesFlag) GetNames() []string {
	return n.names
}

func (n *NamesFlag) String() string {
	return fmt.Sprint(n.names)
}

func (n *NamesFlag) Set(v string) error {
	names := strings.Split(v, ",")
	n.names = append(n.names, names...)

	return nil
}
