package main

import "fmt"

func main() {
	var a = 10;
	b := 10;
	c := 5
	d := 2
	var e int8 = int8(a + b - c * d);

	fmt.Println(e);

	var i = 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}