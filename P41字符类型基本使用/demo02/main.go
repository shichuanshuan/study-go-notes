// 字符串类型基本使用及细节
package main

import "fmt"

func main(){
	// string的基本使用
	var address string = "shichuanshuan"
	fmt.Println("string = ",address)

	// 字符串一旦赋值，字符串就不能修改了：在go中字符串不能修改
	// var str = "hello"
	// str[0] = 'a'  // 这里就不能改变str的内容

	// 字符串另一种表达方法，使用反引号，会省略转义字符
	var arr = `
	var address string = "shichuanshuan"
	fmt.Println("string = ",address)
	`
	fmt.Println("arr = ",arr)
}