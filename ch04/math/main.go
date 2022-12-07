package main

import (
	"fmt"
	"math"
)

func main() {
	base := int(32)
	fmt.Printf("pw: %f\n", math.Pow(float64(base), 2))
}
