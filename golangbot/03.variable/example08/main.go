package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)

	//age := 29      // age是int类型
	//age = "naveen" // 错误，尝试赋值一个字符串给int类型变量
}
