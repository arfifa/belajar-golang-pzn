package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer
	address2.City = "Bogor"

	fmt.Println(address1) // changed to Bogor
	fmt.Println(address2) // changed to Bogor

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}