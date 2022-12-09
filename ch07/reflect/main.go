package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func main() {
	getValueTypeCase()
	fmt.Println("----------------------------------------------------------------")

	complexCase()
	fmt.Println("----------------------------------------------------------------")

	p := &person{
		Id:           uuid.NewString(),
		Name:         "Kut Zhang",
		Age:          44,
		MobileNumber: "1333333333",
	}
	p.GoHome()
	p.Work()
	printMethods(p)
}

// get value type case
func getValueTypeCase() {
	pPerson := &person{
		Id:   uuid.NewString(),
		Name: "Kut Zhang",
	}

	typeRef := reflect.ValueOf(pPerson).Elem()
	valueType := typeRef.Type()
	fmt.Printf("type of pPerson is: %s\n", valueType)
}

// complex case
func complexCase() {
	describeValue := func(value any) {
		refValue := reflect.ValueOf(value).Elem()
		valueType := refValue.Type()

		fmt.Printf("Got type %q, there are %d fields:\n", valueType, valueType.NumField())
		for i := 0; i < valueType.NumField(); i++ {
			fmt.Printf(
				"%s[%q]:%v\n",
				valueType.Field(i).Name,
				valueType.Field(i).Type,
				refValue.Field(i).Interface(),
			)
		}
	}

	pPerson := newPerson("Lana Chang", 40, "18520779938")
	describeValue(pPerson)
}

func printMethods(p *person) {
	refValue := reflect.ValueOf(p)
	refType := refValue.Type()
	fmt.Printf("%v - %v\n", refValue, refType)

	numMethods := refValue.NumMethod()
	for i := 0; i < numMethods; i++ {
		method := refValue.Method(i)
		fmt.Printf("%s --> %s\n", refType.Method(i).Name, method.Type())
	}
}

// person
type person struct {
	Id           string
	Name         string
	Age          uint
	MobileNumber string
}

// new person
func newPerson(name string, age uint, mobileNumber string) *person {
	return &person{
		Id:           uuid.NewString(),
		Name:         name,
		Age:          age,
		MobileNumber: mobileNumber,
	}
}

func (p *person) Work() {
	fmt.Println("start working.......")
	time.Sleep(time.Second)
	fmt.Println("work complete!")
}

func (p *person) GoHome() {
	fmt.Println("start going home.......")
	time.Sleep(time.Second)
	fmt.Println("go home complete!")
}
