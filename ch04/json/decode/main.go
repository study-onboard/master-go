package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/study-onboard/master-go/ch04/json/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <json filename>", os.Args[0])
		os.Exit(1)
	}

	groups := loadFromDataFile(os.Args[1])
	common.PrintGroups(groups)

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(groups)
	bytes := buffer.Bytes()
	jsonString := string(bytes)

	groups = loadFromJsonString(jsonString)
	common.PrintGroups(groups)
}

// load groups from data file
func loadFromDataFile(filename string) []common.WorkGroup {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open json file error: %s\n", err.Error())
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	groups := make([]common.WorkGroup, 0)
	decoder.Decode(&groups)

	return groups
}

// load groups from JSON string
func loadFromJsonString(jsonString string) []common.WorkGroup {
	reader := bytes.NewReader([]byte(jsonString))
	decoder := json.NewDecoder(reader)
	groups := make([]common.WorkGroup, 0)
	decoder.Decode(&groups)

	return groups
}
