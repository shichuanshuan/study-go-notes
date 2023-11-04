// go语言没有 while循环和do while循环
// 可以用for循环实现两种while循环效果
package main

import (
	"fmt"
)

func main(){
	// 用for循环实现 while 效果 hello world
	var i int = 0;

	for {
		if i > 10{ //循环条件
			break
		}
		fmt.Printf("hello world\n")
		i++  // 循环变量迭代
	}

	// 用for循环实现do while效果
	var j int = 0
	for {
		fmt.Printf("hello world~\n")
		j++
		
		if j > 10{
			break
		}
	}
}