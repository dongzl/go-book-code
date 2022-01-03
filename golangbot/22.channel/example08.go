package main

import "fmt"

func digits(number int, dch chan int) {
	if number != 0 {
		digit := number % 10
		dch <- digit
		number /= 10
	}
	close(dch)
}

func calcSquares_a(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes_b(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares_a(number, sqrch)
	go calcCubes_b(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
