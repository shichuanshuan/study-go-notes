package main

import (
	"fmt"
)

// 在go中，函数也是一种数据类型
// 可以赋值给一个变量，改变量就是一个函数类型的变量了。通过该变量可以对函数调用

func getSum(n1 int, n2 int)int{
	return n1 + n2
}

func main(){

	// 给变量赋值函数类型，此刻 a 是函数类型‘可以调用 a 
	a := getSum
	fmt.Printf("a的数据类型%T  getSum的数据类型%T\n",a, getSum )

	// 例如
	res := a(3, 5)
	fmt.Println("res = ",res)
}