package main

import "fmt"

func assert_a(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = 56
	assert_a(s)
	var i interface{} = "Steven Paul"
	assert_a(i)
}
