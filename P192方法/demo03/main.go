package main

import "fmt"

// 1. 声明一个结构体Cirle，字段为radius
// 2. 声明一个方法area和Circle绑定， 可以返回面积
// 3. 提示：画出area执行过程+说明
type Cirle struct{
	radius float64
}

func (cirle Cirle) area()float64{
	return 3.14 * cirle.radius * cirle.radius
}

func main(){
	var c Cirle
	c.radius = 4
	res := c.area()
	fmt.Println("area = ", res)
}