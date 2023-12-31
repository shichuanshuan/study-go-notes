// 字符类型基本使用
package main

import "fmt"

func main(){
	var c1 byte = 'a'

	// 直接输出，输出的会是字符的码值
	fmt.Println("c1 = ",c1)

	// 如果要输出字符，要使用格式化输出
	fmt.Printf("c1 = %c \n",c1)

	// c2 byte = '石'  超出了范围，会溢出
	var c2 int = '石'
	fmt.Printf("c2 = %c",c2)
}