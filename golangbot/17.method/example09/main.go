package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
