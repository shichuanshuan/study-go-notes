// 递归调用案例
package main

import (
	"fmt"
)

func test1(n int){
	if n > 2{    // 2、函数的局部变量是独立的，不会相互影响
		n--
		test1(n) // 3、递归必须向退出递归的条件逼退，否则陷入死循环
	}

	fmt.Printf("n = %d\n", n)
}

// func test2(n int){
// 	if n > 2{
// 		n--
// 		test2(n)
// 	} else {
// 		fmt.Println("n = ",n)
// 	}
// }

func main(){
	test1(4)
}