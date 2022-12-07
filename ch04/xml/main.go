package main

import (
	"encoding/xml"
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

	xmlData, err := xml.MarshalIndent(&groups, "", "    ")
	if err != nil {
		fmt.Printf("xml encode failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(xmlData))

	var target []common.WorkGroup
	err = xml.Unmarshal(xmlData, &target)
	if err != nil {
		fmt.Printf("xml decode failed: %s\n", err.Error())
		os.Exit(1)
	}
	common.PrintGroups(target)
}
