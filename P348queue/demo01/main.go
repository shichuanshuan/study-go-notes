package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用数组实现队列的思路(数组模拟队列)
// 1. 创建一个数组 array， 是作为队列的一个字段
// 2. front 表示队列的头部，初始化为 -1
// 3. rear 表示队列的尾部，初始化为 -1
// 4. 完成队列的基本找查
// AddQueue 加入数据到队列
// GetQueue 从队列中取出数据
// ShowQueue 显示队列

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组模拟队列
	front   int    // 表示指向队列首(不包括首)
	rear    int    // 表示指向队列尾(包括尾)
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if this.rear == this.maxSize-1 { // 重要的提示， rear 是队列尾部(含最后元素)
		return errors.New("queue full")
	}
	this.rear++ // rear 后移
	this.array[this.rear] = val

	return
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// 显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	// this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

// 编写一个主函数测试
func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int

	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			queue.AddQueue(val)
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
