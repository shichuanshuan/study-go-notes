package main

import "fmt"

// 7. Golang接口中不能有任何变量
// type AInterface interface{
// 	Say()
// 	Name string   // 不允许
// }


// 8. 一个接口(比如A接口)可以继承多个别的接口(比如B，C接口)
//    这时如果实现A接口，也必须将B，C接口的方法也全部实现
type BInterface interface{
	test01()
}

type CInterface interface{
	test02()
}

type AInterface interface{
	BInterface
	CInterface
	test03()
}

type Stu struct{

}

func (s Stu) test01(){
	fmt.Println("8. test01...")
}

func (s Stu) test02(){
	fmt.Println("8. test02...")
}

func (s Stu) test03(){
	fmt.Println("8. test03...")
}

func main(){

	// 8. 一个接口(比如A接口)可以继承多个别的接口(比如B，C接口)
	//    这时如果实现A接口，也必须将B，C接口的方法也全部实现
	var stu Stu
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu
	a.test03()
	b.test01()
	c.test02()
	fmt.Println()


	// 9. interface 类型默认是一个指针(引用类型)， 如果没有对
	// interface初始化就使用， 那么会输出 nil


	// 10. 空接口interface{} 没有任何方法，所有类型接口都实现了空接口
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2)
}
