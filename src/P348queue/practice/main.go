package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxsize int
	array   [5]int
	head    int
	tail    int
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("array full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxsize

	return
}

func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxsize == this.head
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("array empty")
	}

	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxsize

	return val, err
}

func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) ListQueue() {
	size := this.Size()

	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxsize
	}
	fmt.Println()

}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxsize - this.head) % this.maxsize
}

func main() {
	queue := &CircleQueue{
		maxsize: 5,
		head:    0,
		tail:    0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入成功")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
