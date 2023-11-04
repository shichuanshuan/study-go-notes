// 加减乘除
package main

import (
	"fmt"
)

// 输入两个数，以及+-*/，并算出结果
func cal(n1 float64, n2 float64, operator byte)float64{
	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Printf("输入有误")
	}

	return res
}

func main(){
	var num1, num2, result float64
	var ope byte

	fmt.Printf("请输入两个数及运算符; 如：a + b\n")

	// 用fmt.Scanln 只能单个地址输入
	fmt.Scanf("%f %c %f",&num1, &ope, &num2)

	fmt.Printf("ope %v\n", ope)
	fmt.Printf("num2 %v\n", num2)

	result = cal(num1, num2, ope)

	fmt.Printf("%v %c %v = %v",num1, ope, num2, result)
	
}