package main

import "fmt"

type HeroNode struct {
	no   int
	name string
	next *HeroNode
}

func InsertHero(head *HeroNode, newHero *HeroNode) {
	// 1.定义一个辅助结点
	temp := head

	for {
		// 2.看看此节点下一个结点是否为空
		if temp.next == nil {
			break
		}

		// 3.若为空，指向下下一个
		temp = temp.next
	}

	// 4.说明此时结点指向的下一个结点为空，给此节点的下一个指向赋予新的结点
	// 也就是链表的末尾添加 结点
	temp.next = newHero
}

func InsertHero2(head *HeroNode, newHero *HeroNode) {
	// 1.定义一个辅助结点
	temp := head
	flag := true

	// 2.对比 id 号
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHero.no {
			// 说明下一个结点的序号大于要插入的序号，
			break
		} else if temp.next.no == newHero.no {
			flag = false
			break
		}

		temp = temp.next
	}

	// 3.判断是否相等
	if !flag {
		fmt.Println("the id same...")
		return
	} else {
		newHero.next = temp.next
		temp.next = newHero
	}
}

func DelHero(head *HeroNode, id int) {
	// 1.定义一个辅助结点
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

	// 2.
	if !flag {
		// 3.
		temp.next = temp.next.next
	}
}

func LisetHero(head *HeroNode) {
	// 定义辅助结点
	temp := head

	// 2.查看此节点的指向是否为空
	if temp.next == nil {
		return
	}

	//
	for {
		fmt.Printf("[%d %s] ==>", temp.next.no, temp.next.name)

		// 3.指向下下一个
		temp = temp.next

		// 4.判断下下一个指向有无
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1.定义结构体
	head := &HeroNode{}

	head1 := &HeroNode{
		no:   1,
		name: "shi",
	}
	head2 := &HeroNode{
		no:   2,
		name: "chu",
	}
	head3 := &HeroNode{
		no:   3,
		name: "sun",
	}

	// 2.加入链表
	// InsertHero(head, head3)
	// InsertHero(head, head1)
	// InsertHero(head, head2)

	// 3.show list
	// LisetHero(head)

	// fmt.Println()
	InsertHero2(head, head3)
	InsertHero2(head, head1)
	InsertHero2(head, head2)
	DelHero(head, 2)

	// 3.show list
	LisetHero(head)
}
