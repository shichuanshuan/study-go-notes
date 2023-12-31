package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的类型 %T, num1的值 %v, num1的地址 %v\n", num1, num1, &num1)

	num2 := new(int)
	// num2的类型 %T = -> *int
	// num2的值      = 地址 0x10c9009c
	// num2的地址 %v = 地址 0x10c84108
	// num2指向的值  = 0
	var ptr *int
	var ptr1 *int
	var ret int = 11
	ptr = &ret
	ptr1 = ptr
	fmt.Printf("ptr %T ptr num = %v\n", ptr1, *ptr)
	// num2 = ptr
	fmt.Printf("num2的类型 %T, num2的值 %v, num2的地址 %v\n, num2指针指向的值%v", num2, num2, &num2, *num2)
}
