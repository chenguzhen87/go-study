//package main
//
//import "fmt"
//
//func main() {
//	a := true
//	b := false
//	fmt.Println("a:", a, "b:", b)
//	c := a && b
//	fmt.Println("c:", c)
//	d := a || b
//	fmt.Println("d:", d)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a int = 89
//	b := 95
//	fmt.Println("value of a is", a, "and b is", b)
//}

//package main
//
//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	var a int = 89
//	b := 95
//	fmt.Println("value of a is", a, "and b is", b)
//	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
//	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	a, b := 5.67, 8.97
//	fmt.Printf("type of a %T b %T\n", a, b)
//	sum := a + b
//	diff := a - b
//	fmt.Println("sum", sum, "diff", diff)
//
//	no1, no2 := 56, 89
//	fmt.Println("sum", no1+no2, "diff", no1-no2)
//}


//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	c1 := complex(5, 7)
//	c2 := 8 + 27i
//	cadd := c1 + c2
//	fmt.Println("sum:", cadd)
//	fmt.Println("c1:", c1)
//	cmul := c1 * c2
//	fmt.Println("product:", cmul)
//}


//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	first := "Naveen"
//	last := "Ramanathan"
//	name := first +" "+ last
//	fmt.Println("My name is",name)
//	fmt.Printf("first type %T",first)
//}


//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	i := 55      //int
//	j := 67.8    //float64
//	sum := i + j //不允许 int + float64
//	fmt.Println(sum)
//}

package main

import (
"fmt"
)

func main() {
	i := 55      //int
	j := 67.8    //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)
	n,err := fmt.Printf("%d,%T",int(j),int(j))
	if err != nil {
		fmt.Println(n,err)
	}

	for true {
		fmt.Println(n,err)
	}

}