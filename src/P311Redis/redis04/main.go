package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 批量操作
func main() {
	// 通过go 向redis 写入数据和读取数据
	// 1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	// 2. 通过go 向redis 写入数据 string [key-val]
	_, err = conn.Do("HMSet", "user02", "name", "Mack", "age", 19)
	if err != nil {
		fmt.Println("set err =", err)
		return
	}

	// 3. 通过go 向redis 读取数据 string [key-val]
	r1, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	// 因为返回 r 是 interface{}
	// 因为 name 对应的值是string, 因此我们需要转换
	// nameString := r.(string)

	for index, value := range r1 {
		fmt.Printf("r[%d]=%s\n", index, value)
	}
}
