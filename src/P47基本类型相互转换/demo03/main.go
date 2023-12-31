package main

import (
	"fmt"
	"strconv"
)

// 基本类型转string类型
func main() {
	var num int8 = 125
	var num2 float64 = 3.1412
	var mychar byte = 's'
	var str string

	// 第一种，使用fmt.Sprintf
	str = fmt.Sprintf("%d", num)
	fmt.Printf("str = %v\n", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str = %v\n", str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str = %v\n", str)

	// 第二种方法,用strconv函数
	str = strconv.FormatInt(int64(num), 10)
	fmt.Printf("str = %v", str)
}