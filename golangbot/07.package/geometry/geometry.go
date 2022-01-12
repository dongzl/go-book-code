package main

import (
	"fmt"
	//"geometry/rectangle" // 导入自定义包
	"log"
)

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	//var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	//fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
