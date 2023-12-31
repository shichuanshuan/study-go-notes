package main

import (
	"fmt"
	"reflect"
)

func reflectTestOne(b interface{}) {

	// 通过反射获取的传入的变量 type， kind， 值
	// 1.先获得 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Printf("1. type =%T\n", rTyp)
	fmt.Println()

	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("2. value =%v type =%T\n", rVal, rVal)
	fmt.Println()

	// 3. 下面将rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("3. iV =%v type =%T\n", iV, iV)

	// 将 interface{} 通过类型断言转成原本得类型
	num2 := iV.(int)

	// 因为rType是接口类型，没有方法；结构体是可以绑定方法的
	// num3 := rTyp.(int)
	// fmt.Println("rTyp = ", num3)
	fmt.Println("3. num2 = ", num2)
}

// 结构体类型
type Student struct {
	Name string
	Age  int
}

func reflectTestTwo(b interface{}) {
	// 1. 结构体类型转成 Type类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("1. Type =", rTyp)

	// 2. 先获得 reflect.ValeOf
	rVal := reflect.ValueOf(b)
	fmt.Printf("2. value =%v\n", rVal)

	// 3.将rVal 转成 interface
	iV := rVal.Interface()

	// 有值可以打印，但是取不出来
	fmt.Printf("3. vale =%v, type =%T\n", iV, iV)

	// 类型断言后，可以把结构体的值取出
	num2 := iV.(Student)
	fmt.Println(num2.Name)
}

func main() {

	// 编写一个案例，
	// 演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	var num = 100
	reflectTestOne(num)
	fmt.Println()

	// 演示结构体数据类型操作
	var Stu Student = Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTestTwo(Stu)

}
