//package main
//
//import "fmt"
//
//func main() {
//	var a chan int
//	if a == nil {
//		fmt.Println("channel a is nil, going to define it")
//		a = make(chan int)
//		fmt.Printf("Type of a is %T", a)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func hello() {
//	fmt.Println("Hello world goroutine")
//}
//func main() {
//	go hello()
//	time.Sleep(1 * time.Second)
//	fmt.Println("main function")
//}


//package main
//
//import (
//	"fmt"
//)
//
//func hello(done chan bool) {
//	fmt.Println("Hello world goroutine")
//	done <- true
//}
//func main() {
//	done := make(chan bool)
//	go hello(done)
//	d := <-done
//	fmt.Printf("main function %t",d)
//}


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func hello(done chan bool) {
//	fmt.Println("hello go routine is going to sleep")
//	time.Sleep(4 * time.Second)
//	fmt.Println("hello go routine awake and going to write to done")
//	done <- true
//}
//func main() {
//	done := make(chan bool)
//	fmt.Println("Main going to call hello go goroutine")
//	go hello(done)
//	<-done
//	fmt.Println("Main received data")
//}


package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)
}