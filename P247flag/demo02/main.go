package main

import "flag"

// 两种常用的定义命令行flag参数的方法
func main() {
	// 方法1
	// flag.Type(flag名，默认值，帮助信息)*Type
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 24, "年龄")
	married := flag.Bool("married", false, "婚否")

	// 方法2
	// flag.TypeVar(Type指针名，flag名，默认值，帮助信息)
	var name2 string
	var age2 int
	var married2 bool
	flag.StringVar(&name2, "name2", "张三", "姓名")
	flag.IntVar(&age2, "age2", 18, "年龄")
	flag.BoolVar(&married2, "married", false, "婚否")
}
