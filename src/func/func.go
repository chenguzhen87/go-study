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

//package main
//
//import (
//	"fmt"
//)
//
//func change(s ...string) {
//	s[0] = "Go"
//	s = append(s, "playground")
//	fmt.Println(s)
//}
//
//func main() {
//	welcome := []string{"hello", "world"}
//	change(welcome...)
//	fmt.Println(welcome)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	a := []int{7, 8, 9}
//	fmt.Printf("%+v\n", a)
//	ap(a)
//	fmt.Printf("%+v\n", a)
//	app(a)
//	fmt.Printf("%+v\n", a)
//
//}
//
//func ap(a []int) {
//	a = append(a, 10)
//}
//
//func app(a []int) {
//	a[0] = 1
//}

//package main
//
//import "fmt"
//
//func f1() (r int) {
//	defer func() {
//		r++
//	}()
//	return 0
//}
//
//func f2() (r int) {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}
//
//func f3() (r int) {
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}
//
//func f4() (int) {
//	var r int = 0
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}
//
//func f5() (r int) {
//
//	// 1.赋值
//	r = 4
//
//	// 2.闭包引用，返回值被修改
//	defer func() {
//		r++
//	}()
//
//	// 3.空的 return
//	return r
//}
//
//
//func main()  {
//	f1 := f1()
//	f2 := f2()
//	f3 := f3()
//	f4 := f4()
//	f5 := f5()
//	fmt.Println(f1)
//	fmt.Println(f2)
//	fmt.Println(f3)
//	fmt.Println(f4)
//	fmt.Println(f5)
//}

//package main
//
//import "fmt"
//
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person.age = 29
//
//
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	s1 := []int{1, 2, 3}
//	s2 := s1[1:]
//	s2[1] = 4
//	fmt.Println(s2)
//	fmt.Println(s1)
//	s2 = append(s2, 5, 6, 7)
//	fmt.Println(s1)
//	fmt.Println(s2)
//}

//package main
//
//import "fmt"
//
//func main() {
//	m := map[int]string{0:"zero",1:"one"}
//	for k,v := range m {
//		fmt.Println(k,v)
//	}
//}

package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}