// 用循环语句打印内容
package main

import (
	"fmt"
)

func main(){
	for i:= 0; i < 10; i++ {
		fmt.Printf("shi\n")
	}

	// 遍历字符串方式1-传统方式
	// 传统方式采用的是字节的方式遍历
	// 而中文，占用三个字节，因此打印不出
	var str string = "abc北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 遍历字符串方式2-for range
	// 对于for range 遍历而言，采用的是字符的方式遍历；
	// 因此，如果字符串有中文，也是能准确打印出来
	str = "shichuanshuan上海"
	for index, val := range str {
		fmt.Printf("index = %d val = %c\n",index, val)
	}
}
