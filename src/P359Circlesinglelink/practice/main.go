package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertNode(head *CatNode, newCat *CatNode) {
	temp := head
	if head.next == nil {
		head.no = newCat.no
		head.name = newCat.name
		head.next = head
		fmt.Println("first node is", newCat)
		return
	}

	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCat
	newCat.next = head
}

func InsertNode2(head *CatNode, newCat *CatNode) *CatNode {
	temp := head
	helper := head
	flag := true
	if head.next == nil {
		head.no = newCat.no
		head.name = newCat.name
		head.next = head
		fmt.Println("first node is", newCat)
		return head
	}

	if temp.next == head {
		if temp.no > newCat.no {
			newCat.next = temp
			temp.next = newCat

			return newCat
		} else {
			temp.next = newCat
			newCat.next = temp

			return head
		}
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	for {
		fmt.Println("temp.next", temp.next)

		if temp.next == head {
			break
		} else if temp.no > newCat.no {
			flag = false

			break
		}
		temp = temp.next
		helper = helper.next
	}
	if !flag {
		newCat.next = temp
		helper.next = newCat
	}
	fmt.Println()

	return head
}

func DelNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("the empty")

		return head
	}

	if temp.next == head {
		temp.next = nil

		return head
	}

	for {
		if helper.next == head {
			break
		}

		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head {
			break
		} else if temp.no == id {
			flag = false
			if temp == head {
				head = head.next
			}
			helper.next = temp.next

			fmt.Println("delete id is ", id)
			break
		}

		temp = temp.next
		helper = helper.next
	}

	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Println("flag delete id is ", id)
		} else {
			fmt.Println("no exsit")
		}
	}

	return head
}

func ListNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		return
	}

	for {
		fmt.Printf("[%d %s]==> ", temp.no, temp.name)

		if temp.next == head {
			break
		}

		temp = temp.next

	}
}

func main() {

	head := &CatNode{}

	head1 := &CatNode{
		no:   1,
		name: "one",
	}
	head2 := &CatNode{
		no:   2,
		name: "two",
	}
	head3 := &CatNode{
		no:   3,
		name: "thress",
	}

	head = InsertNode2(head, head3)
	head = InsertNode2(head, head1)
	head = InsertNode2(head, head2)
	// InsertNode(head, head3)
	// InsertNode(head, head1)
	// InsertNode(head, head2)
	ListNode(head)
	fmt.Println()

	head = DelNode(head, 1)
	head = DelNode(head, 3)
	head = DelNode(head, 2)

	ListNode(head)
}
