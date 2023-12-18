package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.49))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}