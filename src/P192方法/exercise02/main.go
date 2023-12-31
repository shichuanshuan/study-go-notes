package main

import "fmt"

type Num struct{

}

// 4.编写方法：判断一个数是奇数还是偶数
func (num Num) test4(n int){
	if n % 2 == 0{
		fmt.Println("4. 是偶数")
	} else {
		fmt.Println("4. 是奇数")
	}
}


// 5. 根据行、列、字符打印对应行数和列数的字符，
//    比如：行：3， 列：2，字符：*，则打印相应的效果
func (num Num) test5(len int, width int, by string){
	for i := 0; i < len; i++ {
		for j := 0; j < width; j++{
			fmt.Printf(by)
		}
		fmt.Println()
	}
}


// 6. 定义小小计算器结构体（Calcuator）， 实现加减乘除四个功能
//    实现形式1：分四个方法完成；
//    实现方式2：用一个方法搞定
type Calcuator struct{
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) test6(operator byte)float64{
	res := 0.0
	switch operator{
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("Input error")
	}

	return res
}


func main(){
	var n Num
	n.test4(10)
	fmt.Println()


	n.test5(5, 10, "#")
	fmt.Println()


	var c Calcuator
	c.Num1 = 11
	c.Num2 = 6
	res := c.test6('/')
	fmt.Println("6. res =", res)
}