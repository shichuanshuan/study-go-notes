package main

import (
	"fmt"
	"strings"
)

// 1)编写一个函数makeSuffix(suffix string)  可接受一个文件后缀，比如(.jpg) 并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如:jpg),则返回 文件名jpg
// 3）使用闭包的方式完成
// 4)strings.HasSuffix 该函数可以判断某个字符串是否有指定的后缀
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

// 不用闭包实现
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}

	return name
}

func main() {
	f := makeSuffix(".jpg")

	fmt.Println("文件名处理后 ", f("wirte"))
	fmt.Println("文件名处理后 ", f("read.jpg"))

	// 不适用闭包每次都要手动加后缀
	f2 := makeSuffix2(".jpg", "shichuan.jpg")
	fmt.Println("f2 =", f2)
	fmt.Println("f2 =", makeSuffix2(".jpg", "shichuan.jpg"))
}
