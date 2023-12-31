package main

import "fmt"

// 可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效

func test() {

	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数 可以用来捕获异常

		if err != nil {
			fmt.Println("err =", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2

	fmt.Println("res = \n", res)
}

func main() {

	test()

	fmt.Println("main()下的代码...")
}
