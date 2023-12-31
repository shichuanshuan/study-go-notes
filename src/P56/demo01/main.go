// 文件调用
package main

import (
	"fmt"
	"GoCode/P56/demo01H"
)

func main(){

	index := 12
	fmt.Printf("index = %v\n", index)

	// demo01H 是包名 package demo01H
	fmt.Printf("Hare = %v", demo01H.Hare)
}