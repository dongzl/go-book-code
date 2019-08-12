package main

import (
	"fmt"
	"time"
)

func hello(msg string) {
	fmt.Println("Hello " + msg)
}

func main() {
	// 在新的协程中执行 hello 方法
	go hello("World")
	fmt.Println("Run in main")
	// 等待 100 毫秒让协程执行结束
	time.Sleep(100 * time.Microsecond)
}
