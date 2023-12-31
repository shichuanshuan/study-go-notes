package main

import (
	"fmt"
)

func main() {
	// slice的第二种使用方式
	var slice []float64 = make([]float64, 5, 10)
	slice[0] = 10
	slice[1] = 20

	// 对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice 的大小", len(slice))
	fmt.Println("slice 的容量", cap(slice))

	// 方式3，定义一个切片，直接就指定具体数组；
	var strSlice []string = []string{"shi", "chuan", "shuan"}
	fmt.Println("strSlice 的值", strSlice)
	fmt.Println("strSlice 的长度", len(strSlice))
	fmt.Println("strSlice 的容量", cap(strSlice))
}
