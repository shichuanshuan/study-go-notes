package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	arr    [5]int
	top    int
	maxtop int
}

func (this *Stack) Push(val int) error {
	if this.top == this.maxtop-1 {
		return errors.New("push full")
	}

	this.top++
	this.arr[this.top] = val

	return nil
}

func (this *Stack) Pop() (int, error) {
	if this.top == -1 {
		return -1, errors.New("pop empty")
	}

	val := this.arr[this.top]
	this.top--

	return val, nil
}

func (this *Stack) List() {
	if this.top == -1 {
		fmt.Println("list error")
		return
	}

	for i := this.top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 判断是数字还是符号
func (this *Stack) isOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func (this *Stack) isPriority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}

	return res
}

func (this *Stack) Cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	}

	return res
}

func main() {
	numStack := &Stack{
		top:    -1,
		maxtop: 5,
	}

	operStack := &Stack{
		top:    -1,
		maxtop: 5,
	}

	exp := "30+2*6-1-1"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	keepNum := ""

	for {
		ch := exp[index : index+1]

		temp := int([]byte(ch)[0]) // 转为对应的ascii

		if operStack.isOper(temp) {
			if operStack.top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.isPriority(operStack.arr[operStack.top]) >= operStack.isPriority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()

					res := operStack.Cal(num1, num2, oper)

					numStack.Push(res)

					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.isOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		index++
		if index == len(exp) {
			break
		}
	}

	for {
		if operStack.top == -1 {
			break
		}

		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()

		res := operStack.Cal(num1, num2, oper)

		numStack.Push(res)
	}

	res, _ := numStack.Pop()
	fmt.Printf("express =%s = %d\n", exp, res)
}
