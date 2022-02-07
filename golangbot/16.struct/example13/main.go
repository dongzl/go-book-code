package main

import "fmt"

type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}

func main() {
	var spec Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}
