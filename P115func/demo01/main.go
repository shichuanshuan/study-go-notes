// 函数调用
package main

import (
	"fmt"
)

// 实现两数相加、相减
func getSumAndSub(n1 int, n2 int)(sub int, sum int){
	sub = n1 + n2
	sum = n1 - n2

	return sub, sum
}

func main(){
	var result1, result2 int

	result1, result2 = getSumAndSub(2, 4)

	// 希望忽略某个返回值，则使用 _ 符号表示占位忽略
	result3, _ := getSumAndSub(3,4)

	fmt.Printf("之和= %d, 之差 = %d\n",result1, result2)

	fmt.Printf("之和= %d\n",result3)
}