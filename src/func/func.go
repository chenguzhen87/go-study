//package main
//
//import (
//"fmt"
//)
//
//func calculateBill(price, no int) int {
//var totalPrice = price * no
//return totalPrice
//}
//func main() {
//price, no := 90, 6 // 定义 price 和 no,默认类型为 int
//totalPrice := calculateBill(price, no)
//fmt.Println("Total price is", totalPrice) // 打印到控制台上
//}

//package main
//
//import (
//	"fmt"
//)
//
//func rectProps(length, width float64)(float64, float64) {
//	var area = length * width
//	var perimeter = (length + width) * 2
//	return area, perimeter
//}
//
//func main() {
//	area, perimeter := rectProps(10.8, 5.6)
//	fmt.Printf("Area %f Perimeter %f", area, perimeter)
//}

package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}

func test() string {
	var s = `dddddd`
	fmt.Println(s)
	return `1`
}

func main() {
	area, _ := rectProps(10, 5)
	t := test()
	fmt.Println(area, t)
}

