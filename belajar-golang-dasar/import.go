package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	fmt.Println("Bingung Kerja Apa")

	result := helper.SayHello("Rahman")
	fmt.Println(result)
	fmt.Println(helper.Application)
}