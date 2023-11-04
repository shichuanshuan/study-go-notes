package main

import "fmt"

type AInterfance interface{

}

func main(){
	var num float64 = 1.1
	var ainterfance AInterfance
	ainterfance = num  // 空接口可以接收任意类型

	// 类型断言
	// 由于接口是一般类型，不知道具体类型，如果要转成具体类型
	// 就需要使用类型断言，具体的如下：
	y := ainterfance.(float64)
	fmt.Println(y)


	// 类型断言，检查错误
	var num2 float64 = 2.2
	var ainterfance2 AInterfance
	ainterfance2 = num2
	
	y2, ok := ainterfance2.(float64)
	if ok == true{
		fmt.Println(y2)
	} else {
		fmt.Println("error")
	}
}