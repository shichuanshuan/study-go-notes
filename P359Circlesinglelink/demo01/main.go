package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 1.插入链表
func InsertCatNode(head *CatNode, newCat *CatNode) {
	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCat.no
		head.name = newCat.name
		head.next = head
		fmt.Println(newCat, "加入到环形链表")
		return
	}

	// 定义一个临时变量，帮忙，找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}

		temp = temp.next
	}

	// 加入到链表中
	temp.next = newCat
	newCat.next = head
}

// 输出这个环形的链表
func ListCatNode(head *CatNode) {
	fmt.Println("circle link action follow...")
	temp := head
	if temp.next == nil {
		fmt.Println("empty empty")
		return
	}

	for {
		fmt.Printf("cat info=[%d %s]==>  ", temp.no, temp.name)
		if temp.next == head {
			break
		}

		temp = temp.next
	}
}

func main() {
	// 1.定义头节点
	head := &CatNode{}
	head1 := &CatNode{
		no:   1,
		name: "大花猫",
	}

	head2 := &CatNode{
		no:   2,
		name: "小黑猫",
	}

	// 2.增加猫
	InsertCatNode(head, head1)
	InsertCatNode(head, head2)

	// 3.显示
	ListCatNode(head)

}
