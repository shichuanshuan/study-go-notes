package main

import (
	"fmt"
)

var (
	// Fun1就是一个全局匿名函数；全局都可以调用
	// 首字符要大写
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数的时候直接调用，这种方式匿名函数只调用一次

	// 案例演示，两个数之和，使用匿名函数的方式完成
	// 注意没有函数名
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(20, 20)
	fmt.Println("res =", res)

	// 将匿名函数赋给变量 a
	// 则a的数据类型就是函数类型，此时，我们可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(40, 30)
	fmt.Println("res2 =", res2)

	res3 := Fun1(2, 9)
	fmt.Println("res3 =", res3)
}