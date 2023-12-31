package main

import (
	"fmt"
)

func main(){
	// 输入相应字符，显示对应的星期
	// 如：(a, b, c, d, e, f, g)
	// var week byte

	// fmt.Printf("输入字符(a, b, c, d, e, f, g)\n")
	// fmt.Scanf("%c", &week)

	// switch week {
	// 	// case后面可以有多个表达式
	//     case 'a', 'A':
	// 		fmt.Printf("星期一，小猴穿新衣")
	// 	case 'b':
	// 		fmt.Printf("星期二，小猴当小二")
	// 	default:
	// 		fmt.Printf("输入有误")
	// }

	// switch 后面可以不跟表达式和变量，相当于if - else 使用
	var age = 10
	switch {
	case age == 10:
		fmt.Printf("age == 10\n")
	case age == 20:
		fmt.Printf("age == 20\n")
	}

	// case 也可以对范围进行判断
	var socre int = 90
	switch {
	case socre > 90:
		fmt.Printf("优")
	case socre >= 70:
		fmt.Printf("良")
	}
}