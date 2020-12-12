package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int

var mu sync.Mutex // 声明一个互斥锁

var rmu sync.RWMutex // 声明一个读写锁

//存款
func Deposit(amount int) {
	mu.Lock()//获取锁
	balance = balance + amount
	mu.Unlock()//释放锁
}

//读取余额
func Balance() int {
	return balance
}

func main() {
	//小王：存600，并读取余额
	go func() {
		Deposit(600)
		fmt.Println(Balance())
	}()
	//小张：存500
	go func() {
		Deposit(500)
	}()

	time.Sleep(time.Second)
	fmt.Println(balance)
}
