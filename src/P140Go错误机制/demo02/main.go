package main

import (
	"errors"
	"fmt"
)

// 函数去读取以配置init.conf的信息
// 如果文件传输不正确，我们就返回自定义的错误
func readConf(name string) (err error) {
	if name == "config.int" {
		//读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

// 如程序内部的逻辑不一致。任何崩溃都表明了我们的代码中可能存在漏洞，所以对于大部分漏洞，我们应该使用Go语言提供的错误机制，而不是 panic
func test() {

	err := readConf("confi.int")

	if err != nil {
		// 如果读取文件发送错误，就输出这个结果，并终止程序;  main 也会终止
		panic(err)
	}

	fmt.Println("test继续执行")
}

func main() {
	test()

	fmt.Println("main()下面的代码")
}
