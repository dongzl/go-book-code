package main

import "fmt"

func main() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}
