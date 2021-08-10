// https://mp.weixin.qq.com/s/S29FV1oPSAJqCoSwcYpHhA
//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	name     string
//	salary   int
//	currency string
//}
//
///*
//  displaySalary() 方法将 Employee 做为接收器类型
//*/
//func (e Employee) displaySalary() {
//	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
//}
//
//func main() {
//	emp1 := Employee {
//		name:     "Sam Adolf",
//		salary:   5000,
//		currency: "$",
//	}
//	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法
//}
//
//
//package main
//
//import (
//"fmt"
//)
//
//type Employee struct {
//	name     string
//	salary   int
//	currency string
//}
//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	name     string
//	salary   int
//	currency string
//}
//
///*
//displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
//*/
//func displaySalary(e Employee) {
//	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
//}
//
//
//func main() {
//	emp1 := Employee{
//		name:     "Sam Adolf",
//		salary:   5000,
//		currency: "$",
//	}
//	displaySalary(emp1)
//}

//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Rectangle struct {
//	length int
//	width  int
//}
//
//type Circle struct {
//	radius float64
//}
//
//func (r Rectangle) Area() int {
//	return r.length * r.width
//}
//
//func (c Circle) Area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//
//func main() {
//	r := Rectangle{
//		length: 10,
//		width:  5,
//	}
//	fmt.Printf("Area of rectangle %d\n", r.Area())
//	c := Circle{
//		radius: 12,
//	}
//	fmt.Printf("Area of circle %f", c.Area())
//}


//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	name string
//	age  int
//}
//
///*
//使用值接收器的方法。
//*/
//func (e Employee) changeName(newName string) {
//	e.name = newName
//}
//
///*
//使用指针接收器的方法。
//*/
//func (e *Employee) changeAge(newAge int) {
//	e.age = newAge
//}
//
//func main() {
//	e := Employee{
//		name: "Mark Andrew",
//		age:  50,
//	}
//	fmt.Printf("Employee name before change: %s", e.name)
//	e.changeName("Michael Andrew")
//	fmt.Printf("\nEmployee name after change: %s", e.name)
//
//	fmt.Printf("\n\nEmployee age before change: %d", e.age)
//	(&e).changeAge(51)
//	fmt.Printf("\nEmployee age after change: %d", e.age)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	name string
//	age  int
//}
//
///*
//使用值接收器的方法。
//*/
//func (e Employee) changeName(newName string) {
//	e.name = newName
//}
//
///*
//使用指针接收器的方法。
//*/
//func (e *Employee) changeAge(newAge int) {
//	e.age = newAge
//}
//
//func main() {
//	e := Employee{
//		name: "Mark Andrew",
//		age:  50,
//	}
//	fmt.Printf("Employee name before change: %s", e.name)
//	e.changeName("Michael Andrew")
//	fmt.Printf("\nEmployee name after change: %s", e.name)
//
//	fmt.Printf("\n\nEmployee age before change: %d", e.age)
//	e.changeAge(51)
//	fmt.Printf("\nEmployee age after change: %d", e.age)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type address struct {
//	city  string
//	state string
//}
//
//func (a address) fullAddress() {
//	fmt.Printf("Full address: %s, %s", a.city, a.state)
//}
//
//type person struct {
//	firstName string
//	lastName  string
//	address
//}
//
//func main() {
//	p := person{
//		firstName: "Elon",
//		lastName:  "Musk",
//		address: address {
//			city:  "Los Angeles",
//			state: "California",
//		},
//	}
//
//	p.fullAddress() //访问 address 结构体的 fullAddress 方法
//}

//package main
//
//import (
//	"fmt"
//)
//
//type rectangle struct {
//	length int
//	width  int
//}
//
//func area(r rectangle) {
//	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
//}
//
//func (r rectangle) area() {
//	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
//}
//
//func main() {
//	r := rectangle{
//		length: 10,
//		width:  5,
//	}
//	area(r)
//	r.area()
//
//	p := &r
//	/*
//	   compilation error, cannot use p (type *rectangle) as type rectangle
//	   in argument to area
//	*/
//	//area(p)
//
//	p.area()//通过指针调用值接收器
//}


package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}