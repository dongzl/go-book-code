package main

import (
	"fmt"
)

func main() {
	//var a int = 89
	//b := 95
	//fmt.Println("value of a is", a, "and b is", b)
	//fmt.Println("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
	//fmt.Println("type of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
	//
	//c1 := complex(2, 7)
	//c2 := 3 + 8i
	//cadd := c1 + c2
	//fmt.Println("sum: ", cadd)
	//
	//cmul := c1 * c2
	//fmt.Println("product:", cmul)
	//
	//var defaultName = "Sam" // 允许
	//type myString string
	//var customName myString = "Sam" // 允许
	//// customName = defaultName // 不允许
	//customName = myString(defaultName) // 不允许
	//fmt.Println("customName:", customName)
	//
	//if a := 10; a % 2 == 0 {
	//
	//} else {
	//
	//}
	//
	//finger := 4
	//switch finger {
	//case 1:
	//	fmt.Println("Thumb")
	//	fallthrough
	//case 2:
	//	fmt.Println("Index")
	//case 3:
	//	fmt.Println("Middle")
	//case 4:
	//	fmt.Println("Ring")
	//	fallthrough
	////case 4://重复项
	////	fmt.Println("Another Ring")
	//case 5:
	//	fmt.Println("Pinky")
	//}
	//
	//switch num := 75; { // num is not a constant
	//case num < 50:
	//	fmt.Printf("%d is lesser than 50\n", num)
	//	fallthrough
	//case num < 100:
	//	fmt.Printf("%d is lesser than 100\n", num)
	//	fallthrough
	//case num < 200:
	//	fmt.Printf("%d is lesser than 200", num)
	//}

	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6, 7)
	fmt.Println(len(a), cap(a)) // len=7,cap=8
	//const a = 5
	//var intVar int = a
	//var int32Var int32 = a
	//var float64Var float64 = a
	//var complex64Var complex64 = a
	//fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

//func (a int) add(b int) {
//}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}
