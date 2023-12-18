package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func main() {
	result := getHello("Rahman")
	fmt.Println(result)

	fmt.Println(getHello("Eko"))
	fmt.Println(getHello("Blackot"))
}