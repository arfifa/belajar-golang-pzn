package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// alamat1 := new(Address)
	// alamat2 := alamat1

	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}