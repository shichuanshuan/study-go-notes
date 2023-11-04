package main

import "fmt"

// 方法和函数的区别
// 1. 调用方式不一样

// 2. 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递
// 反之亦然

// 3. 对于方法（如struct的方法）， 接收者为类型时，可以直接用指针
// 的变量调用方法，反之亦然
type Box struct{
	len int
	width int
	height int
}

func (box *Box) getVolumn()int{
	return box.len * box.width * (*box).height
}

func main(){
	var b Box

	b.len = 3
	b.width = 3
	b.height = 3

	res := b.getVolumn()

	fmt.Printf("Volumn = %v\n", res)
}