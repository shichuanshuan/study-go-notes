package main

import "fmt"

type DoubleQueue struct {
	no   int
	name string
	pre  *DoubleQueue
	next *DoubleQueue
}

func InsertQueue(head *DoubleQueue, newDouble *DoubleQueue) {
	temp := head

	for {
		if temp.next == nil {
			break
		}

		temp = temp.next
	}
	temp.next = newDouble
	newDouble.pre = temp
}

func InsertQueue2(head *DoubleQueue, newDouble *DoubleQueue) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newDouble.no {
			break
		} else if temp.next.no == newDouble.no {
			flag = false
			break
		}

		temp = temp.next
	}

	if !flag {
		fmt.Println("the id exist")
	} else {
		newDouble.next = temp.next
		newDouble.pre = temp
		if temp.next != nil {
			temp.next.pre = newDouble
		}
		temp.next = newDouble
	}
}

func DelDoubleQueue(head *DoubleQueue, id int) {
	temp := head
	flag := true

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = false
			break
		}

		temp = temp.next
	}

	if !flag {
		fmt.Println("the id exist")
	} else {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	}

}

func ListQueue(head *DoubleQueue) {
	temp := head

	if temp.next == nil {
		return
	}

	for {
		fmt.Printf("[%d %s]==> ", temp.next.no, temp.next.name)

		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListQueue2(head *DoubleQueue) {
	// 逆向显示
	temp := head

	if temp.next == nil {
		return
	}

	for {
		if temp.next == nil {
			break
		}

		temp = temp.next
	}

	for {
		fmt.Printf("[%d %s]==>", temp.no, temp.name)
		temp = temp.pre

		if temp.pre == nil {
			break
		}
	}
}

func main() {
	// 1.创建结构体
	head := &DoubleQueue{}

	head1 := &DoubleQueue{
		no:   1,
		name: "shi",
	}

	head2 := &DoubleQueue{
		no:   2,
		name: "chu",
	}

	head3 := &DoubleQueue{
		no:   3,
		name: "sun",
	}

	// 2.加入到链表
	InsertQueue2(head, head3)
	InsertQueue2(head, head1)
	InsertQueue2(head, head2)

	// 3.显示
	// ListQueue(head)

	// 逆向
	// fmt.Println()
	ListQueue2(head)
}
