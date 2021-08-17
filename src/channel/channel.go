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
//	"time"
//)
//
//func hello(i int) {
//	fmt.Println("Hello world goroutine",i)
//}
//func main() {
//
//	for i := 1; i < 10; i++ {
//		go hello(i)
//	}
//
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


//package main
//
//import (
//	"fmt"
//)
//
//func calcSquares(number int, squareop chan int) {
//	fmt.Println("Final output calcSquares", number)
//	sum := 0
//	for number != 0 {
//		digit := number % 10
//		sum += digit * digit
//		number /= 10
//	}
//	squareop <- sum
//}
//
//func calcCubes(number int, cubeop chan int) {
//	fmt.Println("Final output calcCubes", number)
//	sum := 0
//	for number != 0 {
//		digit := number % 10
//		sum += digit * digit * digit
//		number /= 10
//	}
//	cubeop <- sum
//}
//
//func main() {
//	number := 589
//	sqrch := make(chan int)
//	cubech := make(chan int)
//	go calcSquares(number, sqrch)
//	go calcCubes(number, cubech)
//	squares, cubes := <-sqrch, <-cubech
//	fmt.Println("Final output squares", squares)
//	fmt.Println("Final output cubes", cubes)
//	//fmt.Println("Final output", squares + cubes)
//}


//package main
//
//import (
//	"fmt"
//)
//
//
//func main() {
//	ch := make(chan string, 2)
//	ch <- "naveen"
//	ch <- "paul"
//	fmt.Println(<- ch)
//	fmt.Println(<- ch)
//}



//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func write(ch chan int) {
//	for i := 0; i < 5; i++ {
//		ch <- i
//		fmt.Println("successfully wrote", i, "to ch")
//	}
//	close(ch)
//}
//
//func main() {
//	ch := make(chan int, 2)
//	go write(ch)
//	time.Sleep(2 * time.Second)
//	for v := range ch {
//		fmt.Println("read value", v,"from ch")
//		time.Sleep(2 * time.Second)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func mi(c chan string)  {
//	for a :=  range c {
//		fmt.Println("a",a)
//	}
//
//}
//func main() {
//	ch := make(chan string, 1)
//	go mi(ch)
//	ch <- "naveen"
//	ch <- "paul"
//	ch <- "steve"
//	ch <- "naveen"
//	time.Sleep(time.Second)
//	ch <- "1"
//	ch <- "2"
//	ch <- "3"
//	ch <- "4"
//	ch <- "5"
//	ch <- "6"
//	ch <- "7"
//	ch <- "8"
//	ch <- "9"
//	ch <- "10"
//	ch <- "11"
//	time.Sleep(time.Second)
//	//fmt.Println(<-ch)
//	//fmt.Println(<-ch)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func childgo1(ch chan int)  {
//	for i := 0; i < 10; i++ {
//		ch <- i
//	}
//    close(ch)
//}
//
//func childgo2(ch chan int)  {
//	for v := range ch {
//		fmt.Println("childgo2",v)
//	}
//}
//
//func childgo3(ch chan int)  {
//	for v := range ch {
//		fmt.Println("childgo3",v)
//	}
//}
//
//func main() {
//	ch := make(chan int,4)
//	go childgo1(ch)
//	go childgo2(ch)
//	go childgo3(ch)
//
//	for v := range ch {
//		fmt.Println("main",v)
//	}
//
//	time.Sleep(time.Second)
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func process(i int, wg *sync.WaitGroup) {
//	fmt.Println("started Goroutine ", i)
//	time.Sleep(2 * time.Second)
//	fmt.Printf("Goroutine %d ended\n", i)
//	wg.Done()
//}
//
//func main() {
//	no := 3
//	var wg sync.WaitGroup
//	for i := 0; i < no; i++ {
//		wg.Add(1)
//		go process(i, &wg)
//	}
//	wg.Wait()
//	fmt.Println("All go routines finished executing")
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//type Job struct {
//	id       int
//	randomno int
//}
//type Result struct {
//	job         Job
//	sumofdigits int
//}
//
//var jobs = make(chan Job, 10)
//var results = make(chan Result, 10)
//
//func digits(number int) int {
//	sum := 0
//	no := number
//	for no != 0 {
//		digit := no % 10
//		sum += digit
//		no /= 10
//	}
//	time.Sleep(2 * time.Second)
//	return sum
//}
//func worker(wg *sync.WaitGroup) {
//	for job := range jobs {
//		output := Result{job, digits(job.randomno)}
//		results <- output
//	}
//	wg.Done()
//}
//func createWorkerPool(noOfWorkers int) {
//	var wg sync.WaitGroup
//	for i := 0; i < noOfWorkers; i++ {
//		wg.Add(1)
//		go worker(&wg)
//	}
//	wg.Wait()
//	close(results)
//}
//func allocate(noOfJobs int) {
//	for i := 0; i < noOfJobs; i++ {
//		randomno := rand.Intn(999)
//		job := Job{i, randomno}
//		jobs <- job
//	}
//	close(jobs)
//}
//func result(done chan bool) {
//	for result := range results {
//		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
//	}
//	done <- true
//}
//func main() {
//	startTime := time.Now()
//	noOfJobs := 100
//	go allocate(noOfJobs)
//	done := make(chan bool)
//	go result(done)
//	noOfWorkers := 10
//	createWorkerPool(noOfWorkers)
//	<-done
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time taken ", diff.Seconds(), "seconds")
//}


package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	fmt.Println("end")
}