package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
}
