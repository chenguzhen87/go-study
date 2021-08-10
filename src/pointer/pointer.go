//https://mp.weixin.qq.com/s/lV8Mcq3bX8Zt4BoojtBWuw
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	b := 255
//	var a *int = &b
//	fmt.Printf("Type of a is %T\n", a)
//	fmt.Println("address of b is", a)
//}


//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	a := 25
//	var b *int
//	if b == nil {
//		fmt.Println("b is", b)
//		b = &a
//		fmt.Println("b after initialization is", b)
//	}
//}

//package main
//import (
//	"fmt"
//)
//
//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//	*a++
//	fmt.Println("new value of b is", b)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func change(val *int) {
//	*val = 55
//}
//func main() {
//	a := 58
//	fmt.Println("value of a before function call is",a)
//	b := &a
//	change(b)
//	fmt.Println("value of a after function call is", a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func modify(arr *[3]int) {
//	(*arr)[0] = 90
//}
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(&a)
//	fmt.Println(a)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//func modify(arr *[3]int) {
//	arr[0] = 90
//}
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(&a)
//	fmt.Println(a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func modify(sls []int) {
//	sls[0] = 90
//}
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(a[:])
//	fmt.Println(a)
//}

//package main
//
//func main() {
//	b := [...]int{109, 110, 111}
//	p := &b
//	p++
//}