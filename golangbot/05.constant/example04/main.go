package main

import (
	"fmt"
)

func main() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst // 允许
	fmt.Println("defaultBool = ", defaultBool)
	var customBool myBool = trueConst // 允许
	fmt.Println("customBool = ", customBool)
	//defaultBool = customBool // 不允许
}
