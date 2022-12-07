package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/study-onboard/master-go/ch04/json/common"
)

func main() {
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

	data, err := json.Marshal(&groups)
	if err != nil {
		fmt.Printf("encode json failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("json: %s\n", string(data))
}
