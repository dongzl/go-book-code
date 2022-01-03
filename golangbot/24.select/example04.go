package main

import (
	"fmt"
	"time"
)

func server1_a(ch chan string) {
	ch <- "from server1"
}
func server2_a(ch chan string) {
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1_a(output1)
	go server2_a(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
