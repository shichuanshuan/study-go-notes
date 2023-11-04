package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` // 反引号
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	// 1. 创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	// 2. 将monster变量序列化为 json格式字串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json err", err)
	}

	fmt.Println("jsonStr", jsonStr)
	fmt.Println()

	// 未反序列化前应该是jsonStr {"Name":"牛魔王","Age":500,"Skill":"芭蕉扇"}
	// 反序列化后
	fmt.Println("反序列化后jsonStr", string(jsonStr))
}
