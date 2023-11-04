package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改
// num int 的值
// 修改 student的值
func reflect01(b interface{}) {
	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 看看 rVal的Kind是
	fmt.Printf("rVal kind =%v\n", rVal.Kind())
	// 3. rVal
	// rVal.SetInt(20) // 错误操作，原因是rVal 本身就是地址类型
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10

	// 加地址
	reflect01(&num)
	fmt.Println("num =", num)
}

// 给你一个变量 var v float64 = 1.2，请使用反射来得到它的 reflect.Value,
// 然后获取对应的Type, Kind和值，并将reflect.Value转换成interface{},
// 再将interface{} 转换成float64

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main() {
// 	a := 1
// 	b := 2
// 	defer calc("1", a, calc("10", a, b))
// 	a = 0
// 	defer calc("2", a, calc("20", a, b))
// 	b = 1
// }
