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

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	// 如果只有一个结点
	if (*temp).next == head { // 只有一个结点
		temp.next = nil

		return head
	}

	// 将helper 定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	// 如果有两个包含两个以上结点
	flag := true
	for {
		if temp.next == head {
			// 如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == head { // 说明删除的是头节点
				head = head.next
			}

			// 恭喜找到，我们也可以直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     // 移动 【比较】
		helper = helper.next // 移动 【一旦找到要删除的结点 helper】
	}
	// 这里还有比较一次
	if flag { // 如果flag 为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	}

	return head
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
	head3 := &CatNode{
		no:   3,
		name: "tom",
	}

	// 2.增加猫
	InsertCatNode(head, head1)
	InsertCatNode(head, head2)
	InsertCatNode(head, head3)

	// 3.显示
	ListCatNode(head)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// // 删除
	head = DelCatNode(head, 1)

	ListCatNode(head)

}
