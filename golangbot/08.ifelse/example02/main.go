package main

import "fmt"

func main() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}
