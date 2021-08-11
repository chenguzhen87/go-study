//package main

//type Employee struct {
//	firstName string
//	lastName  string
//	age       int
//}

//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}

//var employee struct {
//	firstName, lastName string
//	age int
//}

//package main
//
//import (
//"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//
//	//creating structure using field names
//	emp1 := Employee{
//		firstName: "Sam",
//		age:       25,
//		salary:    500,
//		lastName:  "Anderson",
//	}
//
//	//creating structure without using field names
//	emp2 := Employee{"Thomas", "Paul", 29, 800}
//
//	fmt.Printf("Employee 1%+v", emp1)
//	fmt.Println("Employee 2", emp2)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	emp3 := struct {
//		firstName, lastName string
//		age, salary         int
//	}{
//		firstName: "Andreah",
//		lastName:  "Nikola",
//		age:       31,
//		salary:    5000,
//	}
//
//	fmt.Printf("Employee 3%+v", emp3)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	var emp4 Employee //zero valued structure
//	fmt.Println("Employee 4", emp4)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	emp5 := Employee{
//		firstName: "John",
//		lastName:  "Paul",
//	}
//	fmt.Println("Employee 5", emp5)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	emp6 := Employee{"Sam", "Anderson", 55, 6000}
//	fmt.Println("First Name:", emp6.firstName)
//	fmt.Println("Last Name:", emp6.lastName)
//	fmt.Println("Age:", emp6.age)
//	fmt.Printf("Salary: $%d", emp6.salary)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	var emp7 Employee
//	emp7.firstName = "Jack"
//	emp7.lastName = "Adams"
//	fmt.Println("Employee 7:", emp7)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
//	fmt.Println("First Name:", (*emp8).firstName)
//	fmt.Println("Age:", (*emp8).age)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//type Employee struct {
//	firstName, lastName string
//	age, salary         int
//}
//
//func main() {
//	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
//	fmt.Println("First Name:", emp8.firstName)
//	fmt.Println("Age:", emp8.age)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Person struct {
//	string
//	int
//}
//
//func main() {
//	p := Person{"Naveen", 50}
//	fmt.Println(p)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Person struct {
//	string
//	int
//}
//
//func main() {
//	var p1 Person
//	p1.string = "naveen"
//	p1.int = 50
//	fmt.Println(p1)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Address struct {
//	city, state string
//}
//type Person struct {
//	name    string
//	age     int
//	address Address
//}
//
//func main() {
//	var p Person
//	p.name = "Naveen"
//	p.age = 50
//	p.address = Address{
//		city:  "Chicago",
//		state: "Illinois",
//	}
//	fmt.Println("Name:", p.name)
//	fmt.Println("Age:", p.age)
//	fmt.Println("City:", p.address.city)
//	fmt.Println("State:", p.address.state)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Address struct {
//	city, state string
//}
//type Person struct {
//	name string
//	age  int
//	Address
//}
//
//func main() {
//	var p Person
//	p.name = "Naveen"
//	p.age = 50
//	p.Address = Address{
//		city:  "Chicago",
//		state: "Illinois",
//	}
//	fmt.Println("Name:", p.name)
//	fmt.Println("Age:", p.age)
//	fmt.Println("City:", p.city) //city is promoted field
//	fmt.Println("State:", p.state) //state is promoted field
//}

//package main
//
//import (
//	"fmt"
//	"go-tudy/src/struct/computer"
//)
//
//func main() {
//	var spec computer.Spec
//	spec.Maker = "apple"
//	spec.Price = 50000
//	fmt.Println("Spec:", spec)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type name struct {
//	firstName string
//	lastName string
//}
//
//
//func main() {
//	name1 := name{"Steve", "Jobs"}
//	name2 := name{"Steve", "Jobs"}
//	if name1 == name2 {
//		fmt.Println("name1 and name2 are equal")
//	} else {
//		fmt.Println("name1 and name2 are not equal")
//	}
//
//	name3 := name{firstName:"Steve", lastName:"Jobs"}
//	name4 := name{}
//	name4.firstName = "Steve"
//	if name3 == name4 {
//		fmt.Println("name3 and name4 are equal")
//	} else {
//		fmt.Println("name3 and name4 are not equal")
//	}
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//type image struct {
//	data map[int]int
//}
//
//func main() {
//	image1 := image{data: map[int]int{
//		0: 155,
//	}}
//	image2 := image{data: map[int]int{
//		0: 155,
//	}}
//	if image1 == image2 {
//		fmt.Println("image1 and image2 are equal")
//	}
//}


//package main
//
//import "go-tudy/src/struct/employee"
//
//func main() {
//	e := employee.Employee {
//		FirstName: "Sam",
//		LastName: "Adolf",
//		TotalLeaves: 30,
//		LeavesTaken: 20,
//	}
//	e.LeavesRemaining()
//}

package main

import "go-tudy/src/struct/employee"

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}