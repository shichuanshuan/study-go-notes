package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Rigth *Hero
}

// 前序遍历[先输出root 结点，然后再输出左子树，最后再输出右子树]
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Rigth)
	}
}

// 中序遍历[先输出root 的左子树结点，然后输出 root 结点，最后再输出 root 的右子树]
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Rigth)
	}
}

// 后序遍历[先输出root 的左子树结点，然后输出 root 的右子树结点，最后再输出 root 结点]
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Rigth)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	// 构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}

	node10 := &Hero{
		No:   10,
		Name: "tom",
	}
	node12 := &Hero{
		No:   12,
		Name: "jack",
	}

	left1.Left = node10
	left1.Rigth = node12

	rigth1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}

	root.Left = left1
	root.Rigth = rigth1

	rigth2 := &Hero{
		No:   4,
		Name: "林冲",
	}

	rigth1.Rigth = rigth2

	// PreOrder(root) // 前序遍历

	// InfixOrder(root) // 中序遍历

	PostOrder(root) // 后序遍历
}
