package main

import "fmt"

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	fmt.Println(map1)
	fmt.Println(map2)
	//if map1 == map2 { //error
	//}
}
