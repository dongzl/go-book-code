package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p.area() //通过指针调用值接收器
}
