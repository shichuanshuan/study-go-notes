package main

import (
	"bufio"
	"fmt"
	"os"
)

// 创建保存棋盘状态的结构体
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子为 1
	chessMap[2][3] = 2 // 白字为 2

	// 2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf(" %d ", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组， 想 -> 算法
	// 思路
	// 1). 遍历 chessMap，如果我们发现有一个元素的值不为 0
	// 创建一个 node 结构体
	// 2). 将其放入到对应的切片即可

	var sparseArr []ValNode

	// 标准的一个稀疏数组应该还有一个， 记录元素的二维数组的规模(行和列，默认值)
	// 创建一个ValNode 值结点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个 ValNode 值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	for index, value := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", index, value.row, value.col, value.val)
	}

	// 存盘
	filePath := "./sparse.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	// 准备写入语句
	str := "shichuan"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	// 把语句写进去
	writer.WriteString(str)

	// 因为write是带缓存，因此在调Writestring方法时，其实
	// 内容是先写道缓存的，所有需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中
	writer.Flush()

	// 打开文件
	file, err = os.Open(filePath)
	if err != nil {
		fmt.Println("Open file err=", err)
	}

	// 当函数退出时，要即使的关闭file
	defer file.Close()

	reader := bufio.NewReader(file)

	str, _ = reader.ReadString('\n') // 读到一个就换行

	// 输出内容
	fmt.Printf(str)

}
