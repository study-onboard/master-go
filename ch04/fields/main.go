package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args

	// check arguments count
	if len(arguments) < 3 {
		fmt.Printf("usage: %s column <file1> [<file2> [.... < fileN]]\n", arguments[0])
		os.Exit(1)
	}

	// validate column
	column, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Printf("Column `%s` is not an integer\n", arguments[1])
		os.Exit(1)
	}
	if column < 0 {
		fmt.Printf("Invalid column number: %d\n", column)
		os.Exit(1)
	}

	// handle file
	for _, filename := range arguments[2:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s: %s\n", filename, err.Error())
			continue
		}
		defer file.Close()
		handleFile(file, column)
	}
}

func handleFile(file *os.File, column int) {
	fmt.Println(file.Name())
	fmt.Println("========================================================================")
	defer fmt.Println("------------------------------------------------------------------------")

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("read content of `%s` failed: %s\n", file.Name(), err.Error())
				return
			}
		}

		columns := strings.Fields(line)
		if len(columns) > column {
			fmt.Println(columns[column])
		}
	}
}
