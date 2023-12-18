package main

import "fmt"

func main() {
		var names [3]string

		names[0] = "Arfifa"
		names[1] = "Rahman"
		names[2] = "Wu"

		fmt.Println(names[0])
		fmt.Println(names[1])
		fmt.Println(names[2])

		// names[3] = "Error"

		var values = [3]int16{95, 90, 87}

		fmt.Println(values)
		fmt.Println(len(values))
}