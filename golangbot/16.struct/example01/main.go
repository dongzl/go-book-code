package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	//creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
