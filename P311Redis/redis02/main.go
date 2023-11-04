package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 集合操作
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	_, err = c.Do("Set", "key1", 998)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "key1"))
	if err != nil {
		fmt.Println("get key1 failed,", err)
		return
	}
	fmt.Println(r)
}
