package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const DATA_FILE_PATH_1 = "./data/1"
const DATA_FILE_PATH_2 = "./data/2"
const DATA_FILE_PATH_3 = "./data/3"

func main() {
	writeStringFileDirectly()
	printLine()
	writeStringByIoutil()
	printLine()
	writeStringByIo()
	printLine()
}

// print line
func printLine() {
	fmt.Println("-------------------------------------------------")
}

// write string to file directly
func writeStringFileDirectly() {
	file, err := os.Create(DATA_FILE_PATH_1)
	if err != nil {
		fmt.Printf("create file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString("Show me the money")
	fmt.Println("write string to file directly complete!")
}

// write string by ioutil
func writeStringByIoutil() {
	ioutil.WriteFile(DATA_FILE_PATH_2, []byte("how do you turn this on"), 0644)
	fmt.Println("write string by ioutil complete!")
}

// write string by io
func writeStringByIo() {
	file, err := os.Create(DATA_FILE_PATH_3)
	if err != nil {
		fmt.Printf("create file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	io.WriteString(file, "Good good study, day day up!!!")
	fmt.Println("write string by io complete!")
}
