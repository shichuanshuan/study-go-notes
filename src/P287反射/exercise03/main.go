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

	if kd != reflect.Ptr && val.Elem().Kind() != reflect.Struct {
		fmt.Println("1. expect struct")
		return
	}

	// 获得该结构体有几个字段
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白象精")

	for i := 0; i < num; i++ {
		fmt.Printf("2. Field %d：值为=%v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("3. struct has %d fields\n", num)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("4. struct has %d methods\n", numOfMethod)
	val.Elem().Method(1).Call(nil)
}

func main() {
	var a Monster = Monster{
		Name:  "黄狮子",
		Age:   408,
		Score: 92.8,
	}
	TestStruct(&a)
}
