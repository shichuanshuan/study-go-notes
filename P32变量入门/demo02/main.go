// 变量的细节
package main

import "fmt"

func main() {

	// 多个变量声明方式
	// 第一种：一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)

	// 第二种：
	var index, index2, name = 10, 20, "tom"
	fmt.Println("index = ", index, "index2 = ", index2, "name = ", name)
}