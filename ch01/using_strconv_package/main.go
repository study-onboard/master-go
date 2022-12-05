package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Please give one or more floats!")
		os.Exit(1)
	}

	arguments := os.Args

	min, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "`%s` is not a float number\n", arguments[1])
		os.Exit(1)
	}

	max := min

	for i := 2; i < len(arguments); i++ {
		number, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "`%s` is not a float number\n", arguments[i])
			os.Exit(1)

		}

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	fmt.Printf("Max float number is: %f\n", max)
	fmt.Printf("Min float number is: %f\n", min)
}
