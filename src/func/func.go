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

//package main
//
//import "fmt"
//
//func rectProps(length, width float64) (area, perimeter float64) {
//	area = length * width
//	perimeter = (length + width) * 2
//	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
//}
//
//func test() string {
//	var s = `dddddd`
//	fmt.Println(s)
//	return `1`
//}
//
//func main() {
//	area, _ := rectProps(10, 5)
//	t := test()
//	fmt.Println(area, t)
//}
//

// https://mp.weixin.qq.com/s/DeEuZcWqE-LzYS_sg9m-vQ
//package main
//
//import (
//	"fmt"
//)
//
//func find(num int, nums ...int) {
//	fmt.Printf("type of nums is %T\n", nums)
//	found := false
//	for i, v := range nums {
//		if v == num {
//			fmt.Println(num, "found at index", i, "in", nums)
//			found = true
//		}
//	}
//	if !found {
//		fmt.Println(num, "not found in ", nums)
//	}
//	fmt.Printf("\n")
//}
//func main() {
//	find(89, 89, 90, 95)
//	find(45, 56, 67, 45, 90, 109)
//	find(78, 38, 56, 98)
//	find(87)
//}


//package main
//
//import (
//	"fmt"
//)
//
//func find(num int, nums... int) {
//	fmt.Printf("type of nums is %T\n", nums)
//	nums[2] = 100
//	//nums = append(nums,200 )
//	//nums[3] = 200 //panic: runtime error: index out of range [3] with length 3
//	found := false
//	for i, v := range nums {
//		if v == num {
//			fmt.Println(num, "found at index", i, "in", nums)
//			found = true
//		}
//	}
//	if !found {
//		fmt.Println(num, "not found in ", nums)
//	}
//	fmt.Printf("\n")
//}
//func main() {
//	nums := []int{89, 90, 95}
//	find(89, nums...)
//	fmt.Println("nums",nums)
//}

package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}