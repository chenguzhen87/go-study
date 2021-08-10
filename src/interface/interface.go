//package main
//
//import (
//	"fmt"
//)
//
////interface definition
//type VowelsFinder interface {
//	FindVowels() []rune
//}
//
//type MyString string
//
////MyString implements VowelsFinder
//func (ms MyString) FindVowels() []rune {
//	var vowels []rune
//	for _, rune := range ms {
//		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
//			vowels = append(vowels, rune)
//		}
//	}
//	return vowels
//}
//
//func main() {
//	name := MyString("Sam Anderson")
//	var v VowelsFinder
//	v = name // possible since MyString implements VowelsFinder
//	fmt.Printf("Vowels are %c", v.FindVowels())
//
//}


//package main
//
//import (
//	"fmt"
//)
//
//type SalaryCalculator interface {
//	CalculateSalary() int
//}
//
//type Permanent struct {
//	empId    int
//	basicpay int
//	pf       int
//}
//
//type Contract struct {
//	empId  int
//	basicpay int
//}
//
////salary of permanent employee is sum of basic pay and pf
//func (p Permanent) CalculateSalary() int {
//	return p.basicpay + p.pf
//}
//
////salary of contract employee is the basic pay alone
//func (c Contract) CalculateSalary() int {
//	return c.basicpay
//}
//
///*
//total expense is calculated by iterating though the SalaryCalculator slice and summing
//the salaries of the individual employees
//*/
//func totalExpense(s []SalaryCalculator) {
//	expense := 0
//	for _, v := range s {
//		expense = expense + v.CalculateSalary()
//	}
//	fmt.Printf("Total Expense Per Month $%d", expense)
//}
//
//func main() {
//	pemp1 := Permanent{1, 5000, 20}
//	pemp2 := Permanent{2, 6000, 30}
//	cemp1 := Contract{3, 3000}
//	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
//	totalExpense(employees)
//}


//package main
//
//import (
//	"fmt"
//)
//
//type Test interface {
//	Tester()
//}
//
//type MyFloat float64
//
//func (m MyFloat) Tester() {
//	fmt.Println(m)
//}
//
//func describe(t Test) {
//	fmt.Printf("Interface type %T value %v\n", t, t)
//}
//
//func main() {
//	var t Test
//	f := MyFloat(89.7)
//	t = f
//	describe(t)
//	t.Tester()
//}

//package main
//
//import (
//	"fmt"
//)
//
//func describe(i interface{}) {
//	fmt.Printf("Type = %T, value = %v\n", i, i)
//}
//
//func main() {
//	s := "Hello World"
//	describe(s)
//	i := 55
//	describe(i)
//	strt := struct {
//		name string
//	}{
//		name: "Naveen R",
//	}
//	describe(strt)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//func assert(i interface{}) {
//	s := i.(int) //get the underlying int value from i
//	fmt.Println(s)
//	fmt.Printf("s%T",i)
//}
//
//func main() {
//	var s interface{} = 56
//	assert(s)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func assert(i interface{}) {
//	v, ok := i.(int)
//	fmt.Println(v, ok)
//}
//func main() {
//	var s interface{} = 56
//	assert(s)
//	var i interface{} = "Steven Paul"
//	assert(i)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//
//func findType(i interface{}) {
//	switch i.(type) {
//	case string:
//		fmt.Printf("I am a string and my value is %s\n", i.(string))
//	case int:
//		fmt.Printf("I am an int and my value is %d\n", i.(int))
//	default:
//		fmt.Printf("Unknown type\n")
//	}
//}
//func main() {
//	findType("Naveen")
//	findType(77)
//	findType(89.98)
//}

//package main
//
//import "fmt"
//
//type Describer interface {
//	Describe()
//}
//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) Describe() {
//	fmt.Printf("%s is %d years old", p.name, p.age)
//}
//
//func findType(i interface{}) {
//	switch v := i.(type) {
//	case Describer:
//		v.Describe()
//	default:
//		fmt.Printf("unknown type\n")
//	}
//}
//
//func main() {
//	findType("Naveen")
//	p := Person{
//		name: "Naveen R",
//		age:  25,
//	}
//	findType(p)
//}
//
//package main
//
//import "fmt"
//
//type Describer interface {
//	Describe()
//}
//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) Describe() { // 使用值接受者实现
//	fmt.Printf("%s is %d years old\n", p.name, p.age)
//}
//
//type Address struct {
//	state   string
//	country string
//}
//
//func (a *Address) Describe() { // 使用指针接受者实现
//	fmt.Printf("State %s Country %s", a.state, a.country)
//}
//
//func main() {
//	var d1 Describer
//	p1 := Person{"Sam", 25}
//	d1 = p1
//	d1.Describe()
//	p2 := Person{"James", 32}
//	d1 = &p2
//	d1.Describe()
//
//	var d2 Describer
//	a := Address{"Washington", "USA"}
//
//	/* 如果下面一行取消注释会导致编译错误：
//	   cannot use a (type Address) as type Describer
//	   in assignment: Address does not implement
//	   Describer (Describe method has pointer
//	   receiver)
//	*/
//	//d2 = a
//
//	d2 = &a // 这是合法的
//	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
//	d2.Describe()
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//type SalaryCalculator interface {
//	DisplaySalary()
//}
//
//type LeaveCalculator interface {
//	CalculateLeavesLeft() int
//}
//
//type Employee struct {
//	firstName string
//	lastName string
//	basicPay int
//	pf int
//	totalLeaves int
//	leavesTaken int
//}
//
//func (e Employee) DisplaySalary() {
//	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
//}
//
//func (e Employee) CalculateLeavesLeft() int {
//	return e.totalLeaves - e.leavesTaken
//}
//
//func main() {
//	e := Employee {
//		firstName: "Naveen",
//		lastName: "Ramanathan",
//		basicPay: 5000,
//		pf: 200,
//		totalLeaves: 30,
//		leavesTaken: 5,
//	}
//	var s SalaryCalculator = e
//	s.DisplaySalary()
//	var l LeaveCalculator = e
//	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
//}

//package main
//
//import (
//	"fmt"
//)
//
//type SalaryCalculator interface {
//	DisplaySalary()
//}
//
//type LeaveCalculator interface {
//	CalculateLeavesLeft() int
//}
//
//type EmployeeOperations interface {
//	SalaryCalculator
//	LeaveCalculator
//}
//
//type Employee struct {
//	firstName string
//	lastName string
//	basicPay int
//	pf int
//	totalLeaves int
//	leavesTaken int
//}
//
//func (e Employee) DisplaySalary() {
//	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
//}
//
//func (e Employee) CalculateLeavesLeft() int {
//	return e.totalLeaves - e.leavesTaken
//}
//
//func main() {
//	e := Employee {
//		firstName: "Naveen",
//		lastName: "Ramanathan",
//		basicPay: 5000,
//		pf: 200,
//		totalLeaves: 30,
//		leavesTaken: 5,
//	}
//	var empOp EmployeeOperations = e
//	empOp.DisplaySalary()
//	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
//}


package main

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
}