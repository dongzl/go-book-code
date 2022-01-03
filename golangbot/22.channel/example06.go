package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	//fmt.Println(<-sendch)

	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}
