//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//	Addr string `json:"addr,omitempty"`
//}
//
//func main() {
//	p1 := Person{
//		Name: "Jack",
//		Age:  22,
//	}
//
//	data1, err := json.Marshal(p1)
//	if err != nil {
//		panic(err)
//	}
//
//	// p1 没有 Addr，就不会打印了
//	fmt.Printf("%s\n", data1)
//
//	// ================
//
//	p2 := Person{
//		Name: "Jack",
//		Age:  22,
//		Addr: "China",
//	}
//
//	data2, err := json.Marshal(p2)
//	if err != nil {
//		panic(err)
//	}
//
//
//	// p2 则会打印所有
//	 fmt.Sprintf("%s\n", data2)
//
//}

//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type Person struct {
//	Name string ``
//	Age string
//}
//
//
//func main() {
//	p := reflect.TypeOf(Person{})
//	name, _ := p.FieldByName("Name")
//	fmt.Printf("%q\n", name.Tag) //输出 ""
//	age, _ := p.FieldByName("Age")
//	fmt.Printf("%q\n", age.Tag) // 输出 ""
//}

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name        string `label:"Name is: "`
	Age         int    `label:"Age is: "`
	Gender      string `label:"Gender is: " default:"unknown"`
}

func main() {
	Print(Person{"chenzhen",20,"1"})
}

func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)
	name := reflect.TypeOf(obj).Name()

	fmt.Println(name)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		fmt.Println("name",field.Name)
		tag := field.Tag
		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")
		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}