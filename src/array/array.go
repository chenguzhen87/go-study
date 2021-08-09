// 数组
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var a [3]int //int array with length 3
//	fmt.Println(a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var a [3]int //int array with length 3
//	a[0] = 12 // array index starts at 0
//	a[1] = 78
//	a[2] = 50
//	fmt.Println(a)
//	for i := 0; i< len(a);i++{
//		fmt.Println(a[i])
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	a := [3]int{12}
//	fmt.Println(a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
//	fmt.Println(a)
//}

//package main
//
//func main() {
//	a := [3]int{5, 78, 8}
//	var b [5]int
//	b = a // not possible since [3]int and [5]int are distinct types
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := [...]string{"USA", "China", "India", "Germany", "France"}
//	b := a // a copy of a is assigned to b
//	b[0] = "Singapore"
//	fmt.Println("a is ", a)
//	fmt.Println("b is ", b)
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := [...]float64{67.7, 89.8, 21, 78}
//	fmt.Println("length of a is",len(a))
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := [...]float64{67.7, 89.8, 21, 78}
//	for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
//		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := [...]float64{67.7, 89.8, 21, 78}
//	sum := float64(0)
//	for i, v := range a {//range returns both the index and value
//		fmt.Printf("%d the element of a is %.2f\n", i, v)
//		sum += v
//	}
//	fmt.Println("\nsum of all elements of a",sum)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func printarray(a [3][2]string) {
//	for _, v1 := range a {
//		for _, v2 := range v1 {
//			fmt.Printf("%s ", v2)
//		}
//		fmt.Printf("\n")
//	}
//}
//
//func main() {
//	a := [3][2]string{
//		{"lion", "tiger"},
//		{"cat", "dog"},
//		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
//	}
//	fmt.Println(a)
//	printarray(a)
//	var b [3][2]string
//	b[0][0] = "apple"
//	b[0][1] = "samsung"
//	b[1][0] = "microsoft"
//	b[1][1] = "google"
//	b[2][0] = "AT&T"
//	b[2][1] = "T-Mobile"
//	fmt.Printf("\n")
//	printarray(b)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	a := [5]int{76, 77, 78, 79, 80}
//	var b []int = a[1:4] // creates a slice from a[1] to a[3]
//	fmt.Println(b)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
//	dslice := darr[2:5]
//	fmt.Println("array before", darr)
//	for i := range dslice {
//		dslice[i]++
//	}
//	fmt.Println("array after", darr)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	numa := [3]int{78, 79 ,80}
//	nums1 := numa[:] // creates a slice which contains all elements of the array
//	nums2 := numa[:]
//	fmt.Println("array before change 1", numa)
//	fmt.Println("nums1", nums1)
//	fmt.Println("nums2", nums2)
//	nums1[0] = 100
//	fmt.Println("array after modification to slice nums1", numa)
//	nums2[1] = 101
//	fmt.Println("array after modification to slice nums2", numa)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
//	fruitslice := fruitarray[1:3]
//	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
//	fruitslice := fruitarray[1:3]
//	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
//	fruitslice = fruitslice[:cap(fruitslice)] // re-slicing furitslice till its capacity
//	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	i := make([]int, 5, 5)
//	fmt.Println(i)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	cars := []string{"Ferrari", "Honda", "Ford"}
//	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
//	cars = append(cars, "Toyota")
//	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
//}