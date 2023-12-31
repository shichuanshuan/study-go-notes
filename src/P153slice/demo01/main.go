package main

import (
	"fmt"
)

func main() {
	// 切片的基本使用
	var intArr [5]int = [...]int{12, 43, 2, 29, 72}

	// 声明/定义一个切片
	// slice 相当于结构体
	/*
		type slice struct{
			ptr *[2]int
			len int
			cap int
		}*/
	// slice := intArr[1:3]
	// 1. slice是切片名
	// 2. intArr[1:3]  表示 slice 引用到intArr这个数组
	// 3. 引用intArr这个数组的起始下标为1，最后下标为3.但是不包括3
	slice := intArr[1:3]

	fmt.Println("slice addr ", &slice[0])
	fmt.Printf("type =%T", slice)
	fmt.Printf("intArr type =%T", intArr)

	fmt.Println("intArr =", intArr)
	fmt.Println("slice 元素是 =", slice)
	fmt.Println("slice 元素的个数 =", len(slice))
	fmt.Println("slice 元素的容量 =", cap(slice)) // 切片的容量是可以动态变化的。容量一般是个数的 两 倍

	fmt.Println("intarr ", &intArr[1])
}
