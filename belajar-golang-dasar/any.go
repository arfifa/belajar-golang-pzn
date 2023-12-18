package main

import "fmt"

func Ups() any  {
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}