package main

import "fmt"

func main() {
	const firstName string = "Arfifa"
	const lastName = "Rahman"

	fmt.Println(firstName, lastName);

	const (
		nickname = "Rahman Wu"	
		age int16 = 28
	)

	fmt.Println(nickname, age)

}