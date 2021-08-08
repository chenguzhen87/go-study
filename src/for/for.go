//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	for i := 1; i <= 10; i++ {
//		fmt.Printf(" %d",i)
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	for i := 1; i <= 10; i++ {
//		if i > 5 {
//			break //loop is terminated if i > 5
//		}
//		fmt.Printf("%d ", i)
//	}
//	fmt.Printf("\nline after for loop")
//}

//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	for i := 1; i <= 10; i++ {
//		if i%2 == 0 {
//			continue
//		}
//		fmt.Printf("%d ", i)
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	i := 0
//	for ;i <= 10; { // initialisation and post are omitted
//		fmt.Printf("%d ", i)
//		i += 2
//	}


//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	i := 0
//	for i <= 10 { //semicolons are ommitted and only condition is present
//		fmt.Printf("%d ", i)
//		i += 2
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
//		fmt.Printf("%d * %d = %d\n", no, i, no*i)
//	}
//
//}

package main

import "fmt"

func main() {
	for {
		fmt.Println("Hello World")
	}
}