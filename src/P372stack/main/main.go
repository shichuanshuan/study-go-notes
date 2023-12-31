package main

import (
	"errors"
	"fmt"
)

// 创建一个数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    // 表示我们栈最大可以存放数个数
	Top    int    // 表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr    [5]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	// 存放数据
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}

	// 先取值， 再 this.Top--
	val = this.arr[this.Top]
	this.Top--

	return val, nil
}

func (this *Stack) List() {
	// 先判断是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("the stack status follow")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}

	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// 显示
	stack.List()

	// 出栈
	val, _ := stack.Pop()
	fmt.Println("Pop stack val =", val)

	// 显示
	stack.List()
}
