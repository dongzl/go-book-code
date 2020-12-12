package main

import "fmt"

func main() {
	var a interface{} = 10
	t, ok := a.(int)
	t = 11
	fmt.Println(t)
	fmt.Println(a)
	if ok {
		fmt.Println("int", t)
	}
}

//func main()  {
//	msg := make(chan string)
//	go func() { msg <- "from goroutine"}()
//	fmt.Println(<-msg)
//}

//Go 并发
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	go say("world")
//	say("hello")
//}

//错误处理
//func Sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, errors.New("math: square root of negative number")
//	}
//	return math.Sqrt(f), nil
//}
//
//func main() {
//	result, err := Sqrt(-1)
//	if err != nil {
//		fmt.Println(err)
//		fmt.Println(result)
//	}
//	result1, _ := Sqrt(2)
//	fmt.Println(result1)
//}
