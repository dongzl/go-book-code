package main

import (
	"fmt"
	"sync"
)

var z = 0

func increment_c(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment_c(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of z", z)
}
