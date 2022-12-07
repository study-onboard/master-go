package common

import "fmt"

// print groups
func PrintGroups(groups []WorkGroup) {
	for _, group := range groups {
		fmt.Printf("----%-6s---------------------------\n", group.Name)
		for _, member := range group.Members {
			fmt.Printf("%s - %s\n", member.Name, member.Position)
		}
		fmt.Println()
		fmt.Println()
	}
}
