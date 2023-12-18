package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Just copy value not refferance
	// address2.City = "Bogor"

	// fmt.Println(address1) // not change
	// fmt.Println(address2) // changed to Bogor

	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer refferance
	address2.City = "Bogor"

	fmt.Println(address1) // changed to Bogor
	fmt.Println(address2) // changed to Bogor
}