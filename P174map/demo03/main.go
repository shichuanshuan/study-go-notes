package main

import "fmt"

// map 的增删改查
func main() {
	var cities = make(map[string]string)

	// 没有就是增加，已有的就是修改
	cities["No.1"] = "jack"
	cities["No.2"] = "Make"
	cities["No.3"] = "Dai"

	fmt.Println("cities =", cities)
	fmt.Println()

	// 删除操作，调用内置函数delete
	// 不支持一次性删除所有
	delete(cities, "No.3")
	fmt.Println("cities delete =", cities)
	fmt.Println()

	// 查,如果有这个值，则返回ture
	index, BOOL := cities["No.1"]
	if BOOL {
		fmt.Println("Find num value =", index)
	} else {
		fmt.Println("No find num")
	}
	fmt.Println()

	// 如果想要一次性删除所有
	// 1. 循环删除
	// 2. 直接给要删除的变量 make 一个新空间
	cities = make(map[string]string)
	fmt.Println("cities =", cities)
}
