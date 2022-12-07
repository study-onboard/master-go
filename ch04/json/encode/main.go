package main

import (
	"bufio"
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

	groups := []common.WorkGroup{
		{
			Name: "Masaga",
			Members: []common.Member{
				{
					Name:     "Kut Zhang",
					Position: "IT Manager",
				},
				{
					Name:     "Lana Chang",
					Position: "Boss",
				},
			},
		},
		{
			Name: "Lusico",
			Members: []common.Member{
				{
					Name:     "Lucy Liu",
					Position: "CEO",
				},
				{
					Name:     "Lily Pu",
					Position: "CTO",
				},
			},
		},
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(groups)
	fmt.Printf("json: %s\n", string(buffer.Bytes()))

	file, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		fmt.Printf("open file %s failed: %s\n", os.Args[1], err.Error())
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		fmt.Printf("write json file failed: %s\n", err.Error())
		os.Exit(1)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Printf("flush json file failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("write json file complete.")
}
