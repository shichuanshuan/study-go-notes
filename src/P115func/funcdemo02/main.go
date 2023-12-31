package main

import (
	"fmt"
)

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，在go中，函数可以做形参，被调用
// func(int, int)int是类型，  和 int 一样是类型
func myFun(funvar func(int, int) int, n1 int, n2 int) int {
	return funvar(n1, n2)
}

func main() {
	res2 := myFun(getSum, 20, 30)
	fmt.Println("res2= ", res2)
}
