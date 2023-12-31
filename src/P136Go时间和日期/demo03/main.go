package main

import (
	"fmt"
	"strconv"
	"time"
)

// 获取一个程序执行的时间
func test() {
	str := " "

	for i := 0; i < 1000; i++ {

		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	strat := time.Now().Unix()

	test()

	end := time.Now().Unix()

	fmt.Printf("执行test()耗费时间为%v秒\n", end-strat)
}
