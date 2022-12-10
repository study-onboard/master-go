package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("get docker client failed: %s\n", err.Error())
		os.Exit(1)
	}

	listContainers(cli)
	printLine()
	listImages(cli)
}

func printLine() {
	fmt.Println("----------------------------------------------------------------")
}

// list containers
func listContainers(cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Printf("list containers failed: %s\n", err.Error())
		os.Exit(1)
	}
	for i, container := range containers {
		fmt.Printf("%d. %v\n", i, container.Names)
	}
}

// list images
func listImages(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		fmt.Printf("list images failed: %s\n", err.Error())
		os.Exit(1)
	}

	for i, image := range images {
		fmt.Printf("%d. %s\n", i, image.RepoTags[0])
	}
}
