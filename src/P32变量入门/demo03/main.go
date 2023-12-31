// golang中+的运算
package main

import "fmt"

func main(){
	// 整型做加法运算
	var num1, num2 = 10, 20
	var sum = 0
	sum = num1 + num2
	fmt.Println("sum = ", sum)

	// 字符型做拼接
	var str1, str2 = "shi", "chuan"
	var res = str1 + str2
	fmt.Println("res = ",res)
}