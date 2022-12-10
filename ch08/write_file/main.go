package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const DATA_FILE_PATH_1 = "./data/1"
const DATA_FILE_PATH_2 = "./data/2"
const DATA_FILE_PATH_3 = "./data/3"
const DATA_FILE_PATH_4 = "./data/4"

func main() {
	writeStringFileDirectly()
	printLine()
	writeStringByIoutil()
	printLine()
	writeStringByIo()
	printLine()
	writeStringByBufio()
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

// write string by bufio
func writeStringByBufio() {
	writeFile, err := os.Create(DATA_FILE_PATH_4)
	if err != nil {
		fmt.Printf("create file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)
	writer.WriteString("show me the money\n")
	writer.WriteString("how do you turn this on\n")
	writer.WriteString("good good study, day day up\n")
	writer.WriteString("make love, no war\n")
	err = writer.Flush()
	if err != nil {
		fmt.Printf("flush writer error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("write string by bufio complete!")

	// ------------------------------------------------

	readFile, err := os.Open(DATA_FILE_PATH_4)
	if err != nil {
		fmt.Printf("open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("read file error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
