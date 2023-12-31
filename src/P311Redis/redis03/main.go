package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// hash表
func main() {
	// 通过go 向redis 写入数据和读取数据
	// 1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	// 2. 通过go 向redis 写入数据 string [key-val]
	_, err = conn.Do("HSet", "user01", "name", "jon")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 3. 通过go 向redis 读取数据 string [key-val]
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	// 因为返回 r 是 interface{}
	// 因为 name 对应的值是string, 因此我们需要转换
	// nameString := r.(string)

	fmt.Printf("操作ok r1=%v r2=%v \n", r1, r2)
}
