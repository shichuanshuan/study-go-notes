// 切片的细节
// 对切片进行切片
package main

import "fmt"

func main() {
	var intArr = [5]int{10, 20, 30, 40, 50}

	slice := intArr[1:4]
	fmt.Println("slice =", slice)

	slice2 := slice[1:3]
	slice2[0] = 100 // 因为slice、slice2和intArr指向的数据空间是同一个，因此slice2 = 100，slice和intArr也会改变

	fmt.Println("slice2 =", slice2)
	fmt.Println("slice =", slice)
	fmt.Println("intArr =", intArr)
}
