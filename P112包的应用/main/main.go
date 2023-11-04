// 包的快速入门
package main

import (
	"fmt"
	"GoCode/P112包的应用/utils"
)

func main(){
	var num1, num2, result float64
	var ope byte

	fmt.Printf("请输入两个数及运算符; 如：a + b\n")

	// 用fmt.Scanln 只能单个地址输入
	fmt.Scanf("%f %c %f",&num1, &ope, &num2)

	result = utils.Cal(num1, num2, ope)

	fmt.Printf("%v %c %v = %v",num1, ope, num2, result)
	
}