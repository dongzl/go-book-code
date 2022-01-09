package main

import (
	"fmt"
)

func main() {
	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam" // 允许
	fmt.Println("customName = ", customName)
	fmt.Println("defaultName = ", defaultName)
	//customName = defaultName // 不允许
}
