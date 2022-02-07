package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p := Person{"Naveen", 50}
	fmt.Println(p)
}
