package main

import "fmt"

type Address struct {
	city, state string
}
type Person struct {
	name    string
	age     int
	address Address
}

func main() {
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}
