package main

import "fmt"

// 课下练习
// 1. 在MethodUtils结构体编个方法，从键盘接收整数（1-9），打印对应乘法表：

// 2. 编写方法，使给定的一个二维数组（3 * 3）转置：
//    行列之间转换
/*
1 2 3      1 4 7
4 5 6      2 5 8
7 8 9      3 6 9
*/

type MethodUtils struct {
}

func (methodutils MethodUtils) test(len int) {

	for num1 := 1; num1 <= len; num1++ {
		for num2 := 1; num2 <= num1; num2++ {

			fmt.Printf("%v * %v = %v  ", num2, num1, num1*num2)

		}
		fmt.Println()
	}
}

func main() {
	var m MethodUtils

	m.test(9)
}
