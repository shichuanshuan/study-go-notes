package main

import (
	"fmt"
)

// 13go支持可变参数
func sum(n1 int, args ...int) int { // args只能在末尾
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

func main() {
	result := sum(10, -1, 2, 5, 6)
	fmt.Println("result =", result)
}