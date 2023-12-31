// 为什么可以动态扩展
package main

// 是个引用类型

import "fmt"

func main(){

	// 定义方式一
	// 基本语法 var 变量名 map[keytype][valuetype]
	var a map[string]string

	// 使用map 要先make make的作用是给map分配空间
	// 分配10个 map[][] 大小的空间
	a = make(map[string]string, 5)
	a["no1"] = "shi"
	a["no2"] = "chuan"
	a["no3"] = "shuan"

	fmt.Println(a)
	fmt.Println("len =", len(a))   // 为什么可以动态扩展


	// 定义方式二
	var num1 map[string]string = make(map[string]string)
	num1["name1"] = "上海"
	num1["name2"] = "深圳"
	num1["name3"] = "北京"
	fmt.Println("num1 =", num1)


	// 定义方式三
	var num2 map[string]string = map[string]string{
		"one" : "shi",
		"two" : "chuan",
	}
	num2["three"] = "shuan"
	fmt.Println("num2 =", num2)
}
