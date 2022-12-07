package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/study-onboard/master-go/ch04/json/common"
)

func main() {
	const FILE_NAME string = "../common/data.json"
	file, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Printf("open json file error: %s\n", err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read json file error: %s\n", err.Error())
		os.Exit(1)
	}

	var groups []common.WorkGroup
	err = json.Unmarshal(data, &groups)
	if err != nil {
		fmt.Printf("decord json failed: %s\n", err.Error())
		os.Exit(1)
	}

	common.PrintGroups(groups)
}
