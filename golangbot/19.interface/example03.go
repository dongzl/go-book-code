package main

import "fmt"

type Describer_a interface {
	Describe()
}

func main() {
	var d1 Describer_a
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}

	d1.Describe()
}
