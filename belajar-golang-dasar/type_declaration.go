package main

import "fmt"

func main() {
	
	type NoKTP string

	var ktpRahman NoKTP = "111111111"

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRahman)
	fmt.Println(contohKtp)
}