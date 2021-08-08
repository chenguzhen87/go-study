//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	finger := 4
//	switch finger {
//	case 1:
//		fmt.Println("Thumb")
//	case 2:
//		fmt.Println("Index")
//	case 3:
//		fmt.Println("Middle")
//	case 4:
//		fmt.Println("Ring")
//	case 5:
//		fmt.Println("Pinky")
//
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	finger := 4
//	switch finger {
//	case 1:
//		fmt.Println("Thumb")
//	case 2:
//		fmt.Println("Index")
//	case 3:
//		fmt.Println("Middle")
//	case 4:
//		fmt.Println("Ring")
//	case 4://重复项
//		fmt.Println("Another Ring")
//	case 5:
//		fmt.Println("Pinky")
//
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	switch finger := 8; finger {
//	case 1:
//		fmt.Println("Thumb")
//	case 2:
//		fmt.Println("Index")
//	case 3:
//		fmt.Println("Middle")
//	case 4:
//		fmt.Println("Ring")
//	case 5:
//		fmt.Println("Pinky")
//	default: // 默认情况
//		fmt.Println("incorrect finger number")
//	}
//}


//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	letter := "i"
//	switch letter {
//	case "a", "e", "i", "o", "u": // 一个选项多个表达式
//		fmt.Println("vowel")
//	default:
//		fmt.Println("not a vowel")
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	num := 75
//	switch { // 表达式被省略了
//	case num >= 0 && num <= 50:
//		fmt.Println("num is greater than 0 and less than 50")
//	case num >= 51 && num <= 100:
//		fmt.Println("num is greater than 51 and less than 100")
//	case num >= 101:
//		fmt.Println("num is greater than 100")
//	}
//
//}

package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {

	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}