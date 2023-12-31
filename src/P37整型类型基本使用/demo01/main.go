// 整型的类型
package main

import "fmt"

func main(){
	// rune 占用空间与 int32等价
	var index rune = 65566
	fmt.Println("index =",index)

	// byte 等价 uint8，取值范围0~255
	var b byte = 255
	fmt.Println("byte = %T", b)
}