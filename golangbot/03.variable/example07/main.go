package main

import "fmt"

func main() {
	// 简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的。
	a, b := 20, 30 // 声明变量a和b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("changed b is", b, "c is", c)

	//a, b := 20, 30 // 声明a和b
	//fmt.Println("a is", a, "b is", b)
	//a, b := 40, 50 // 错误，没有尚未声明的变量
}
