package main

import "fmt"

type slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside
}
