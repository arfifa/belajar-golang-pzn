package main

import "fmt"

func main() {
	// Cara membuat map yang 1 
	// var person map[string]string = map[string]string{}
	// person["name"] = "Arfifa"
	// person["address"] = "Bekasi"

	// Cara lain membuat map
	person := map[string]string{
		"nama" : "Arfifa",
		"address": "Bekasi",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println(person["salah"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Eko Kurniawan Khannedy"
	book["ups"] = "salah"

	fmt.Println(book);

	delete(book, "ups")

	fmt.Println(book)
}