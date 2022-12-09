package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

const DATA_FILE = "./persons.data"

func main() {
	writeData()
	readData()
}

func writeData() {
	persons := []Person{
		{
			Name:    "Kut Zhang",
			Age:     42,
			Comment: "IT Manager",
		},
		{
			Name:    "LuLu Ci",
			Age:     25,
			Comment: "Boss",
		},
	}

	writeFile, err := os.Create(DATA_FILE)
	if err != nil {
		fmt.Printf("create data file error: %s\n", err.Error())
		os.Exit(1)
	}

	encoder := gob.NewEncoder(writeFile)
	err = encoder.Encode(persons)
	if err != nil {
		fmt.Printf("encode data error: %s\n", err.Error())
		os.Exit(1)
	}
}

func readData() {
	var persons []Person
	readFile, err := os.Open(DATA_FILE)
	if err != nil {
		fmt.Printf("open data file error: %s\n", err.Error())
		os.Exit(1)
	}

	decoder := gob.NewDecoder(readFile)
	decoder.Decode(&persons)

	for _, person := range persons {
		fmt.Printf("Name: %s, Age: %d, Comment: %s\n", person.Name, person.Age, person.Comment)
	}
}

type Person struct {
	Name    string
	Age     int
	Comment string
}
