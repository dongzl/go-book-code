package main

import "fmt"

func main()  {
	s := []int{5, 6, 7, 8, 9}
	//fmt.Println(remove(s, 2))
	fmt.Println(remove2(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i + 1:])
	return slice[:len(slice) - 1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
