// 范围大->范围小,比如 int64 转成 int8[-128~127] 编译时不会报错
package main

import (
	"fmt"
)

func main(){
	var num int64 = 999999

	var temp int8 = int8(num)

	// 转换结果溢出 temp != 99999,而是63
	fmt.Printf("temp = %v,num = %v\n",temp, num)
}