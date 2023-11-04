package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 = make(map[int]int)

	map1[4] = 66
	map1[2] = 22
	map1[1] = 10
	map1[6] = 60

	fmt.Printf("map1 = %v\n", map1)

	// 如果按照map的key的顺序进行排序输出
	// 1. 先将map的key 放入到切片中
	// 2. 对切片排序
	// 3. 遍历切片，然后按照key来输出map的值

	var slice []int

	// 为什么要追加，不能fmt.Printf("slice 下标为 %v\n", slice[k])
	// range 目标函数有索引 index 的值就是目标函数的索引，不会从0开始
	// 例如： index = 2, 4, 1, 6
	for index, value := range map1 {
		slice = append(slice, index)
		fmt.Printf("index %v value = %v\n", index, value)
	}
	fmt.Println()

	// 排序
	sort.Ints(slice)
	fmt.Printf("键排序后顺序为 %v\n", slice)
	fmt.Println()

	// 目标函数没有索引
	// index = 0, 1, 2, 3
	for index, k := range slice {
		fmt.Printf("map1[%v] = %v ", k, map1[k])
		fmt.Println("index = ", index)
	}
}
