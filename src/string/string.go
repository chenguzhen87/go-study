//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	name := "Hello World"
//	fmt.Println(name)
//}

// package main

// import (
// 	"fmt"
// )

// func printBytes(s string) {
// 	for i:= 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	printBytes(name)
// }

//package main
//
//import (
//    "fmt"
//)
//
//func printBytes(s string) {
//    for i:= 0; i < len(s); i++ {
//        fmt.Printf("%x ", s[i])
//    }
//}
//
//
//func printChars(s string) {
//    for i:= 0; i < len(s); i++ {
//        fmt.Printf("%c ",s[i])
//    }
//}
//
//func main() {
//    name := "Hello World"
//    printBytes(name)
//    fmt.Printf("\n")
//    printChars(name)
//}


//package main
//
//import (
//	"fmt"
//)
//
//func printBytes(s string) {
//	for i:= 0; i < len(s); i++ {
//		fmt.Printf("%x ", s[i])
//	}
//}
//
//func printChars(s string) {
//	runes := []rune(s)
//	for i:= 0; i < len(runes); i++ {
//		fmt.Printf("%c ",runes[i])
//	}
//}
//
//func main() {
//	name := "Hello World"
//	printBytes(name)
//	fmt.Printf("\n")
//	printChars(name)
//	fmt.Printf("\n\n")
//	name = "Señor"
//	printBytes(name)
//	fmt.Printf("\n")
//	printChars(name)
//}


//package main
//
//import (
//	"fmt"
//)
//
//func printCharsAndBytes(s string) {
//	for index, rune := range s {
//		fmt.Printf("%c starts at byte %d\n", rune, index)
//	}
//}
//
//func main() {
//	name := "Señor"
//	printCharsAndBytes(name)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
//	str := string(byteSlice)
//	fmt.Println(str)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
//	str := string(byteSlice)
//	fmt.Println(str)
//}

package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}