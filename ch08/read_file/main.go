package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
)

const DATA_FILE_PATH = "./data.txt"
const BINARY_FILE_PATH = "./data.bin"
const CSV_FILE_PATH = "./data.csv"

func main() {
	lineByLine()
	printLine()

	wordByWord()
	printLine()

	charByChar()
	printLine()

	readRandomFile()
	printLine()

	writeBinaryFile()
	printLine()

	readBinaryFile()
	printLine()

	readCSVFile()
	printLine()
}

// print line
func printLine() {
	fmt.Println("--------------------------------------------------------")
}

// read file data line by line
func lineByLine() {
	file, err := os.Open(DATA_FILE_PATH)
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Read file error: %s\n", err.Error())
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}
}

// read file data word by word
func wordByWord() {
	// open file
	file, err := os.Open(DATA_FILE_PATH)
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	// handle line
	handleLine := func(line string) {
		pattern := regexp.MustCompile(`[^\s]+`)
		words := pattern.FindAllString(line, -1)
		for _, word := range words {
			fmt.Println(word)
		}
	}

	// read lines
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Read file error: %s\n", err.Error())
				os.Exit(1)
			}
		}
		handleLine(line)
	}
}

// read file data char by char
func charByChar() {
	// open file
	file, err := os.Open(DATA_FILE_PATH)
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	// handle line
	handleLine := func(line string) {
		for _, character := range line {
			fmt.Println(string(character))
		}
	}

	// read lines
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Read file error: %s\n", err.Error())
				os.Exit(1)
			}
		}
		handleLine(line)
	}

}

// read file /dev/random for get random number
func readRandomFile() {
	file, err := os.Open("/dev/random")
	if err != nil {
		fmt.Printf("open file error: %s\n", err.Error())
		os.Exit(1)
	}

	var result int64
	binary.Read(file, binary.LittleEndian, &result)
	fmt.Printf("Got a random number: %d\n", result)
}

// read binary file
func readBinaryFile() {
	// open file
	file, err := os.Open(BINARY_FILE_PATH)
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	// read data by size
	readBytes := func(reader io.Reader, size uint32) []byte {
		var buffer []byte = make([]byte, size)
		_, err := reader.Read(buffer)
		if err != nil {
			fmt.Printf("read file error: %s\n", err.Error())
			os.Exit(1)
		}
		return buffer
	}

	// length, string data, int data
	data := readBytes(file, 4)
	length := binary.BigEndian.Uint32(data)
	data = readBytes(file, length)
	message := string(data)
	fmt.Printf("message: %s\n", message)

	data = readBytes(file, 4)
	dataBuffer := bytes.NewReader(data)
	var subfix int32
	binary.Read(dataBuffer, binary.BigEndian, &subfix)
	fmt.Printf("subfix: %d\n", subfix)
}

// read binary file
func writeBinaryFile() {
	// open file
	file, err := os.OpenFile(BINARY_FILE_PATH, os.O_CREATE|os.O_WRONLY, fs.FileMode(0644))
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	message := "show me the money"
	length := uint32(len(message))
	subfix := int32(128)

	err = binary.Write(file, binary.BigEndian, &length)
	if err != nil {
		fmt.Printf("write file error: %s\n", err.Error())
		os.Exit(1)
	}

	_, err = file.WriteString(message)
	if err != nil {
		fmt.Printf("write file error: %s\n", err.Error())
		os.Exit(1)
	}

	err = binary.Write(file, binary.BigEndian, &subfix)
	if err != nil {
		fmt.Printf("write file error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Write binary file complete!")
}

// read csv file
func readCSVFile() {
	file, err := os.Open(CSV_FILE_PATH)
	if err != nil {
		fmt.Printf("open file error: %s\n", err.Error())
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("read csv file error: %s\n", err.Error())
		os.Exit(1)
	}

	for _, record := range records {
		fmt.Printf("%10s%5s\n", record[0], record[1])
	}
}
