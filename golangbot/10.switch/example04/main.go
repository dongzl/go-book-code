package main

import "fmt"

func main() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}
