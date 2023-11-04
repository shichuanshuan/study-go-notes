package main

import (
	"fmt"
	"os"
)

func main() {

	// 打开文件
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("Open file err=", err)
	}

	// 输出文件，看看文件是什么
	fmt.Printf("file=%v\n", *file)

	// 关闭文件
	err = file.Close()

}
