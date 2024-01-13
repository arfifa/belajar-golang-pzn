package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func MultipleParamater[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParamater[string, int]("eko", 100)
	MultipleParamater[int, string](100, "Rahman")
}
