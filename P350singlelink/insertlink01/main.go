package main

import "fmt"

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 这个表示指向下一个结点
}

// 给链表插入一个结点
// 编写第一种插入方法，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1. 先找到该链表的最后这个结点
	// 2. 创建一个辅助结点 [跑龙套，帮忙]
	temp := head
	for {
		if temp.next == nil { // 表示找到
			break
		}

		temp = temp.next // 让 temp 不断的指向下一个结点
	}

	// 3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode

}

func ListHeroNode(head *HeroNode) {
	// 1. 创建一个辅助结点 [跑龙套，帮忙]
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空 空 空 空!!!")
		return
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d %s %s]==>",
			temp.next.no, temp.next.name, temp.next.nickname)

		// 指向下一个链表
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1.先创建一个头结点
	head := &HeroNode{}

	// 2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	// 3. 加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)

	// 4. 显示
	ListHeroNode(head)
}
