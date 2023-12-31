package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 链接池
// 说明：通过golang 对redis 操作，还可以通过redis 链接池
// 1) 事先初始化一定数量的链接，放入到链接池
// 2) 当Go 需要操作redis时，直接从redis 链接池取出链接jike
// 3) 这样可以节省临时获取redis 链接的时间，从而提高效率

// 定义一个全局的pool
var pool *redis.Pool

// 当启动程序时，就初始化链接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数，0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 先从 pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "cat")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	// 取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r = ", r)
}
