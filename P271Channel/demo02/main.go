package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 定义一个存放任意数据类型的管道
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"

	cat := Cat{"小花猫", 4}
	allChan <- cat

	// 我们希望获得管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	// 下面写法是错误的
	// newCat = <- allChan
	// fmt.Printf("newCat = %v", newCat.Name)

	// 正确写法
	newCat := <-allChan
	a := newCat.(Cat)
	fmt.Printf("newCat.Name = %v", a.Name)
}