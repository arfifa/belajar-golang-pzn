package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "eko" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Message)
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Message)
		// } else {
		// 	fmt.Println("unknown error")
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}