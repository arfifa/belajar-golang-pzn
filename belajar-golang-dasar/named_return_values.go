package main

import "fmt"

func getCompleteNamed() (firstName, middleName, lastName string) {
	firstName = "Arfifa"
	middleName = "Rahman"
	lastName = "Blackot"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteNamed()

	fmt.Println(a, b, c)
}