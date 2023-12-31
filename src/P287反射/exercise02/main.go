package main

import (
	"fmt"
	"reflect"
)

// 定义一个Monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float64
	Sex   string
}

// 方法1，显示 s 的值
func (s Monster) Print() {
	fmt.Println("1. ---start---")
	fmt.Println(s)
	fmt.Println("2. ---end---")
}

// 方法0，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 方法2，接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("3. expect struct")
		return
	}

	// 获得该结构体有几个字段
	num := val.NumField()
	fmt.Printf("4. struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("5. Field %d：值为=%v\n", i, val.Field(i))
		// 获取到struct标签，注意需要通过reflect.Type来获得tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("6. Field %d: tag为=%v\n", i, tagVal)
		}
	}

	// 获得该结构体有几个方法，方法顺序按照字母ACSll排序的
	numOfMethod := val.NumMethod()
	fmt.Printf("7. struct has %d methods\n", numOfMethod)
	val.Method(1).Call(nil)

	// 调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)     // 传入的参数是 []reflect.Value
	fmt.Println("8. res =", res[0].Int()) // 返回结果，返回的结果是 []reflect.Value
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}
