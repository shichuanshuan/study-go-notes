package main

import "fmt"

func main() {
	var i int = 100

	// 将i转换成float类型
	var f float32 = float32(i)

	// 两种输出等价
	fmt.Printf("i = %v f =%v\n", i, f)
	fmt.Println("i =",i, "f =", f)

}
