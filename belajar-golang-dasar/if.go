package main

import "fmt"

func main() {
	name := "Arfifa"

	if name == "Arfifa" {
		fmt.Println("Hello Arfifa")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang!")
	} else {
		fmt.Println("Nama sudah benar.")
	}
}