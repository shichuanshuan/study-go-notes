package main

import (
	"fmt"
)

// 累加器
// 闭包就是一个函数和与其相关的环境组合的一个整体
func AddUpper() func(int) int {
	index := 10

	return func(x int) int {
		index += x

		return index
	}
}

func main() {

	// 使用前面的代码
	f := AddUpper()

	fmt.Println("f =", f(1)) // 11
	fmt.Println("f =", f(2)) // 13
}
