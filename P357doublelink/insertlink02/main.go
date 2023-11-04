package main

import "fmt"

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 这个表示指向上一个结点
	next     *HeroNode // 这个表示指向下一个结点
}

// 给双向链表插入一个结点
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
	newHeroNode.pre = temp
}

// 给链表有序插入一个结点
// 编写第一种插入方法，在单链表的最后加入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1. 先找到该链表的最后这个结点
	// 2. 创建一个辅助结点 [跑龙套，帮忙]
	temp := head
	flag := true

	// 让插入的结点的 no， 和 temp 的下一个结点的 no 比较
	for {
		if temp.next == nil { // 说明到结点的最后
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明 newHeroNode 就应该插入到 temp 后面
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next // 让 temp 不断的指向下一个结点
	}

	if !flag {
		fmt.Println("对不起, 已经存在no =", newHeroNode)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// 删除一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false

	// 让插入的结点的 no， 和 temp 的下一个结点的 no 比较
	for {
		if temp.next == nil { // 说明到结点的最后
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next // 让 temp 不断的指向下一个结点
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry, 要删除的id 不存在")
	}
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
func ListHeroNode2(head *HeroNode) {
	// 1. 创建一个辅助结点 [跑龙套，帮忙]
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空 空 空 空!!!")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d %s %s]==>",
			temp.no, temp.name, temp.nickname)

		// 判断是否有链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	// 1.先创建一个头结点
	head := &HeroNode{}

	// 2. 创建一个新的 HeroNode
	herd1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	herd2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	herd3 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	// 3. 加入
	InsertHeroNode(head, herd3)
	InsertHeroNode(head, herd1)
	InsertHeroNode(head, herd2)
	DelHeroNode(head, 3)

	// 4. 显示
	ListHeroNode(head)

	// fmt.Println()
	// fmt.Println("逆向显示")
	// DelHeroNode(head, 2)
	// ListHeroNode2(head)
}
