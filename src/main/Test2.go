package main

func main() {
	// 创建一个无缓冲的 channel
	ch := make(chan int)
	// 主协程会阻塞在此处，发生死锁
	<- ch
}

