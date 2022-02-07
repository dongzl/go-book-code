package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)
}
