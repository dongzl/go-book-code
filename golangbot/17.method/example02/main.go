package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp1)
}
