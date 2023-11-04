package main

import (
	"fmt"
	"os"
)

type Emp struct {
	id   int
	name string
	next *Emp
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {

	temp := this.Head
	var pre *Emp = nil

	if this.Head == nil {
		this.Head = emp
		return
	}

	for {
		if temp != nil {
			if temp.id > emp.id {
				break
			}
			pre = temp
			temp = temp.next
		} else {
			break
		}
	}

	pre.next = emp
	emp.next = temp
}

func (this *EmpLink) Print(index int) {
	temp := this.Head
	if this.Head == nil {
		fmt.Printf("链表%d 为空", index)
		return
	}

	for {
		fmt.Printf("链表%d 雇员id=%d 名字=%s", index, temp.id, temp.name)

		temp = temp.next
		if temp == nil {
			break
		}
	}
	fmt.Println()
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (hash *HashTable) Insert(emp *Emp) {
	node := hash.HashFun((*emp).id)
	hash.LinkArr[node].Insert(emp)
}

func (hash *HashTable) HashFun(id int) int {
	return id % 7
}

func (hash *HashTable) List() {
	for i := 0; i < len(hash.LinkArr); i++ {
		hash.LinkArr[i].Print(i)
		fmt.Println()
	}
}

func main() {
	sel := ""
	var hash HashTable

	for {
		fmt.Println("============== 请输入你的选择 =================")
		fmt.Println("input 增加雇员信息")
		fmt.Println("show 显示雇员信息")
		fmt.Println("find 查找雇员信息")
		fmt.Println("exit 退出雇员信息")
		fmt.Scanln(&sel)

		switch sel {
		case "input":
			id := 0
			name := ""
			fmt.Println("雇员id")
			fmt.Scanln(&id)
			fmt.Println("雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				id:   id,
				name: name,
			}
			hash.Insert(emp)
		case "show":
			hash.List()
		case "find":
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("没有雇员信息")

		}
	}
}
