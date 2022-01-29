package main

import "fmt"

func change(val *int) {
	*val = 55
}

func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
}
