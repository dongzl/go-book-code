package main

import "fmt"

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
