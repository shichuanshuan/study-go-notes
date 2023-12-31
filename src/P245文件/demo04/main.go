package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 创建一个新文件，写入内容
	// 1.打开文件
	filePath := "d:/ipconfig.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 及时关闭
	defer file.Close()

	// 准备写入语句
	str := "石传栓\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	// 把语句写进去
	writer.WriteString(str)

	// 因为write是带缓存，因此在调Writestring方法时，其实
	// 内容是先写道缓存的，所有需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中
	writer.Flush()

}
