//package main
//
//func main() {
//	const a = 55 // 允许
//	a = 89       // 不允许重新赋值
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println("Hello, playground")
//	var a = math.Sqrt(4)   // 允许
//	const b = math.Sqrt(4) // 不允许
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var name = "Sam"
//	fmt.Printf("type %T value %v", name, name)
//
//}

//package main
//
//func main() {
//	var defaultName = "Sam" // 允许
//	type myString string
//	var customName myString = "Sam" // 允许
//	customName = defaultName // 不允许
//
//}

//package main
//
//func main() {
//	const trueConst = true
//	type myBool bool
//	var defaultBool = trueConst // 允许
//	var customBool myBool = trueConst // 允许
//	defaultBool = customBool // 不允许
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	const a = 5
//	var intVar int = a
//	var int32Var int32 = a
//	var float64Var float64 = a
//	var complex64Var complex64 = a
//	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
//}



//package main
//
//import (
//"fmt"
//)
//
//func main() {
//var i = 5
//var f = 5.6
//var c = 5 + 6i
//fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
//
//}


package main

import (
	"fmt"
)

func main() {
	var a = 5.9/8
	fmt.Printf("a's type %T value %v",a, a)
}