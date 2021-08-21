//package main
//
//import "fmt"
//// 声明单个变量
//func main() {
//	var age int // 变量声明
//	fmt.Println("my age is", age)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var age int // 变量声明
//	fmt.Println("my age is", age)
//	age = 29 // 赋值
//	fmt.Println("my age is", age)
//	age = 54 // 赋值
//	fmt.Println("my new age is", age)
//}


//package main
//
//import "fmt"
//// 声明变量并初始化
//func main() {
//	var age int = 29 // 声明变量并初始化
//
//	fmt.Println("my age is", age)
//}

//package main
//
//import "fmt"
//
//// 类型推断
//func main() {
//	var age = 29 // 可以推断类型
//	fmt.Println("my age is", age)
//}


//package main
//
//import "fmt"
//
//// 声明多个变量
//func main() {
//	var width, height int = 100, 50 // 声明多个变量
//
//	fmt.Println("width is", width, "height is", height)
//}


//package main
//
//import "fmt"
//
//func main() {
//	var width, height int
//	fmt.Println("width is", width, "height is", height)
//	width = 100
//	height = 50
//	fmt.Println("new width is", width, "new height is ", height)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var (
//		name   = "naveen"
//		age    = 29
//		height int
//	)
//
//	fmt.Println("my name is", name, ", age is", age, "and height is", height)
//}


//package main
//
//import "fmt"
//
//func main() {
//	name, age := "naveen", 29 // 简短声明
//
//	fmt.Println("my name is", name, "age is", age)
//}

//package main
//
//import "fmt"
//
//func main() {
//	name, age := "naveen" //error，age 没有被赋值
//
//	fmt.Println("my name is", name, "age is", age)
//}

//package main
//
//import "fmt"
//
//func main() {
//	a, b := 20, 30 // 声明变量a和b
//	fmt.Println("a is", a, "b is", b)
//	b, c := 40, 50 // b已经声明，但c尚未声明
//	fmt.Println("b is", b, "c is", c)
//	b, c = 80, 90 // 给已经声明的变量b和c赋新值
//	fmt.Println("changed b is", b, "c is", c)
//}


//package main
//
//import "fmt"
//
//func main() {
//	a, b := 20, 30 // 声明a和b
//	fmt.Println("a is", a, "b is", b)
//	a, b := 40, 50 // 错误，没有尚未声明的变量
//}

//
//package main
//
//func main() {
//	age := 29      // age是int类型
//	age = "naveen" // 错误，尝试赋值一个字符串给int类型变量
//}


//package main
//
//import "fmt"
//// :=声明变量的方式应该只能在局部作用域里面，不能在全局声明
//age := 29      // age是int类型
//func main() {
//
//	fmt.Println(age)
//}


//package main
//
//import "fmt"
//
//func  f1(a int)  {
//	a = 4
//   fmt.Println(a)
//}
//
//func main() {
//	var a = 0
//	f1(a)
//	fmt.Println(a)
//
//}


//package main
//
//import "fmt"
//
//func  f1(a *int)  {
//	*a = 4
//	fmt.Println(a)
//}
//
//func main() {
//	var a = 0
//	f1(&a)
//	fmt.Println(a)
//
//}

//
//package main
//
//import "fmt"
//
//type per struct {
//	name string
//	age int
//}
//
//func  f1(p *per)  {
//	*p = per{"limei",40}
//	fmt.Println(p)
//}
//
//func main() {
//	var p = per{"chenzhen",30}
//	f1(&p)
//	fmt.Println(p)
//}


package main

import "fmt"

type per struct {
	name string
	age int
}

func  f1(p *per)  {
	p.age = 40
	fmt.Println(p)
}

func main() {
	var p = per{"chenzhen",30}
	f1(&p)
	fmt.Println(p)
}
