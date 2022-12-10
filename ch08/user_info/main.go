package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Printf("current user id - method 1: %d\n", os.Getuid())
	fmt.Printf("current group id - method 1: %d\n", os.Getgid())

	user, err := user.Current()
	if err != nil {
		fmt.Printf("get current user info failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("current user id - method 2: %s\n", user.Uid)
	fmt.Printf("current group id - method 2: %s\n", user.Gid)
}
