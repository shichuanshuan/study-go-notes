package main

import (
	"fmt"
)

func main(){
	// 给int取别名，在go中取别名 myInit 和 int 虽然都是 int 类型，
	// 但go认为 myInit 和 int 是两个类型
	type myInit int

	var num1 myInit
	var num2 int
	num1 = 40
	// 会报错
	// num2 = num1

	// 修改后
	num2 = int(num1)
	fmt.Printf("num1 = %d  num2 = %d\n",num1, num2)
}