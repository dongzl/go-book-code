package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	area, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f ", area)
}
