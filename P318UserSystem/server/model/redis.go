package model

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的pool
var Pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeOut time.Duration) {

	Pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: idleTimeOut, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接代码，连接哪个ip 的 redis
			return redis.Dial("tcp", address)
		},
	}
}
