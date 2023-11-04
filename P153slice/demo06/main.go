package main

import "fmt"

func main(){
	// string 底层是一个byte数组，因此string也可以进行切片处理
	// 输出 shichuanshuan
	var str string = "helloshichuanshuan"
	slice := str[5:]
	fmt.Println("slice =", slice)

	// string是不可变的，也就是说不能通过 str[0] = 'z' 方式修改字符串
	// str[0] = 'z' [编译不会通过，报错，原因是 string 是不可变]

	// 如果要修改字符串，可以先将string -> []byte -> 修改 -> 重写转成 string
	// helloshichuanshuan 修改为 zello...
	arr := []byte(str)
	arr[0] = 'z'
	str = string(arr)
	fmt.Println("str =", str)
	
	
	// 转成 []byte 后， 可以处理英文和数字，但是不能处理中文
	// 原因 []byte 以字节来处理，而一个汉字是3个字节，因此就会出现乱码
	// 解决方法 将 string 转成 []rune 即可，因为 []rune 是按字符来处理，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '上'
	str = string(arr1)
	fmt.Println("str = ", str)
}