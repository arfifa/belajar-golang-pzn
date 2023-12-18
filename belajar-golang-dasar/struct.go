package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var rahman Customer
	fmt.Println(rahman)

	rahman.Name = "Arfifa Rahman"
	rahman.Address = "Indonesia"
	rahman.Age = 28
	fmt.Println(rahman)

	joko := Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 25,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 23}
	fmt.Println(budi)	

	budi.sayHello("Blackot")
	rahman.sayHello("Blackot")
}