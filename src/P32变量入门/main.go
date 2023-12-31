//变量的简介和快速入门
package main

import "fmt"

func main() {

	// 变量的使用方法
	// 第一种:指定变量类型，声明后若不赋值，使用默认值
	// int 的默认值是0
	var i int
	fmt.Println("i = ", i)

	// 第二种：根据值自行判定变量类型（类型推导）
	var num = 10.10
	fmt.Println("num = ", num)

	// 第三种：省略var，注意 := 左侧的变量不是声明过的，否则会导致错误
	// 下面方式等价 var name string name = "shi"
	name := "shi"
	fmt.Println("num = ", name)



	// 多个变量声明方式
	// 第一种：一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)

	// 第二种：
	// var index, 
}
