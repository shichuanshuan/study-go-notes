// 包的函数定义文件
package utils

import (
	"fmt"
)

func Cal(n1 float64, n2 float64, operator byte)float64{
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