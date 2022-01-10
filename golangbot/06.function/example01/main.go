package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func main() {
	price, no := 90, 6 // 定义 price 和 no,默认类型为 int
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice) // 打印到控制台上
}
