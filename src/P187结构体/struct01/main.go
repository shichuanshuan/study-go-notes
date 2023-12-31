package main

import "fmt"

// 结构体定义的四种方式
type Person struct{
	name string
	age int
}
func main(){
	// 方式1
	var p1 Person
	fmt.Println("p1 =", p1)
	fmt.Println()


	// 方式2
	var p2 Person = Person{"Mary", 24}
	fmt.Println("p2 =", p2)
	fmt.Println()


	// 方式3
	// 因为p3 是一个指针，因此标准的给字段赋值方式
	// (*p3).name = "smith" 也可以这样写 p3.name = "smith"
	var p3 *Person = new(Person)
	(*p3).name = "smith"
	(*p3).age = 20
	// *p3.name = "smith"
	// *p3.age = 20
	fmt.Println("p3 =", *p3)
	fmt.Println()


	// 方式4
	var p4 *Person = &Person{"scoft", 2}
	(*p4).name = "scoft"
	(*p4).age = 10
	fmt.Println("p4 =", *p4)
	fmt.Println()
}