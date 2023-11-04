package main

import (
	"fmt"
	"time"
)

/*
go协程的特点
有独立的栈空间
共享程序堆空间
调度由用户控制
协程是轻量级的线程
*/

// 快速入门
// 1) 在主线程中，开启一个goroutine，该协程每隔 1 秒输出"hello, world"
// 2) 在主线程中也每隔一秒输出"hello,golang",输出10次后，退出程序
// 3) 要求主线程和goroutine同时执行
// 4) 画出主线程和协程执行流程图

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("test hello world %d\n", i)
		time.Sleep(time.Second)
	}

}

func main() {
	go test()

	for i := 1; i <= 10; i++ {
		fmt.Printf("main() hello golang %d\n", i)
		time.Sleep(time.Second)
	}
}
