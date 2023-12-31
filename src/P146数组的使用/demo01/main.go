package main

import "fmt"

func main() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 =", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02 =", numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03 =", numArr03)

	// 类型推导
	var numArr04 = [...]string{1: "chuan", 0: "shi", 2: "shuan"}
	fmt.Println("numArr04 =", numArr04)

}
