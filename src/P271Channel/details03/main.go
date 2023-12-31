package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	// 这里我们可以使用defer + recover
	defer func() {
		// 捕获test抛出的 panic
		if err := recover(); err != nil {
			fmt.Println("tese() 发生错误", err)
		}
	}()

	// 定义了一个未make 的 map
	var myMap map[int]string
	myMap[0] = "golang" // error
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second)
	}
}
