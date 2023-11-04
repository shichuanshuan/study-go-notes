// 如何查看某个变量在内存中占用字节大小和数据类型
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var index = "shichuanshuan"

	fmt.Printf("内存大小为:%d\n数据类型为:%T\n", unsafe.Sizeof(index), index)
}
