package main

import "fmt"

type image struct {
	data map[int]int
}

func main() {
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	fmt.Println(image1)
	fmt.Println(image2)
	//if image1 == image2 {
	//	fmt.Println("image1 and image2 are equal")
	//}
}
