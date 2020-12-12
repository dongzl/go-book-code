package main

func main() {
	ch := make(chan int)
	select {
	case i := <- ch:
		println(i)
	default:
		println("default")
	}
}
