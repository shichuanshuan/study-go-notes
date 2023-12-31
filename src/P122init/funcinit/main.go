package main

import (
	"fmt"
)

// 如果一个文件同时包含全局变量定义、init函数 main函数，
// 则执行流程是变量定义->init函数->main函数
var age = test()

func test()int{
	fmt.Println("test()...")

	return 90
}

// 先执行init函数再执行main
// init函数，通常可以在init中完成初始化工作
func init(){
	fmt.Println("init()...")
}

func main(){
	fmt.Println("main()...",age)
}