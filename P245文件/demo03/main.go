package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// 使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/test.txt"
	// 文件内容少时，可使用该方法
	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	// 把读取到的内容显示到终端
	fmt.Printf("%v\n", content)
	fmt.Printf("%v", string(content))
}
