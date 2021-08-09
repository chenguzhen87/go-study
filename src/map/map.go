// https://mp.weixin.qq.com/s/9ql797x43pqGf-xaOTIWiQ
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var personSalary map[string]int
//	if personSalary == nil {
//		fmt.Println("map is nil. Going to make one.")
//		personSalary = make(map[string]int)
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int {
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("personSalary map contents:", personSalary)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	employee := "jamie"
//	fmt.Println("Salary of", employee, "is", personSalary[employee])
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	employee := "jamie"
//	fmt.Println("Salary of", employee, "is", personSalary[employee])
//	fmt.Println("Salary of joe is", personSalary["joe"])
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	newEmp := "joe"
//	value, ok := personSalary[newEmp]
//	if ok == true {
//		fmt.Println("Salary of", newEmp, "is", value)
//	} else {
//		fmt.Println(newEmp,"not found")
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("All items of a map")
//	for key, value := range personSalary {
//		fmt.Printf("personSalary[%s] = %d\n", key, value)
//	}
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("map before deletion", personSalary)
//	delete(personSalary, "steve")
//	fmt.Println("map after deletion", personSalary)
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("length is", len(personSalary))
//
//}


//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("Original person salary", personSalary)
//	newPersonSalary := personSalary
//	newPersonSalary["mike"] = 18000
//	fmt.Println("Person salary changed", personSalary)
//}

//
//package main
//
//func main() {
//	map1 := map[string]int{
//		"one": 1,
//		"two": 2,
//	}
//
//	map2 := map1
//	// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 ni
//	if map1 == map2 {
//	}
//}