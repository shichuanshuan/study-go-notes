// 切片细节
// 对切片进行动态追加
package main

import (
	"fmt"
)

func main() {

	// 用append内置函数, 可以对切片进行动态追加
	var slice []int = []int{11, 22, 33}

	// 通过append直接给result追加具体元素
	var result = make([]int, 2)
	result = append(slice, 44, 55)

	fmt.Println("result =", result)
	fmt.Println("result sizeof", len(result))
	fmt.Println("result cap", cap(result))

	fmt.Println()
	// 通过append将切片slice追加给slice3
	slice2 := append(slice, slice...)
	fmt.Println("slice2 =", slice2)

	// 切片使用copy内置函数完成拷贝
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("a =%T\n", a)

	var slice3 = make([]int, 10)
	fmt.Println("slice3 =", slice3)

	copy(slice3, a)
	fmt.Println("copy after slice3 =", slice3)
}
