package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 打开文件
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("Open file err=", err)
	}

	// 当函数退出时，要即使的关闭file
	defer file.Close()

	// 创建一个 *Reader, 是带缓冲的
	/*
		const (
			defaultBufSize = 4096 // 默认缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)

	// 循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个就换行

		if err == io.EOF {
			break
		}

		// 输出内容
		fmt.Printf(str)
	}
	fmt.Println("read end...")
}
