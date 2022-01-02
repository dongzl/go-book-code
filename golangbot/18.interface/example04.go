package main

import "fmt"

func describe_a(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe_a(s)
	i := 55
	describe_a(i)
	str := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe_a(str)
}
