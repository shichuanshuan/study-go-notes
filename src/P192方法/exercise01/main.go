package main

import "fmt"

// 1. 编写结构体(MethodUtils),编程一个方法，方法不需要参数，
//    在方法中打印一 10 * 8 的矩形，在main方法中调用该方法。
type MethodUtils struct{

}

func (methodutils MethodUtils) test1(){
	for i := 0; i < 10; i++{
		for j := 0; j < 8; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 2. 编写一个方法， 提供m和n两个参数，方法中打印一个 m*n 的矩形
func (methodutils MethodUtils)test2(n int, m int){
	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 3. 编写一个方法算该矩形的面积（可以接收长len， 宽width），
//    将其作为方法返回值，在main方法中调用该方法，接收返回的面积值并打印。
func (methodutils MethodUtils) test3(len float64, width float64)float64{
	return len * width
}

func main(){
	var m MethodUtils
	m.test1()
	fmt.Println()


	m.test2(10, 10)
	fmt.Println()


	var res float64
	res = m.test3(10.1, 20.2)
	fmt.Printf("3. res = %.2f\n", res)
}