package main

import (
	"fmt"
	"sync"
)

var y = 0

func increment_a(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment_a(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}
