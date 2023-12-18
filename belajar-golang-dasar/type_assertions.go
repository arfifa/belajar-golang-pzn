package main

import "fmt"

func random() any {
	return false
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	
	// error assertion
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int: 
		fmt.Println("Int", value)
	default: 
		fmt.Println("Unknown", value)
	}
}