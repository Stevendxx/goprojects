package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var connPool *redis.Pool

func init() {
	connPool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 180,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := connPool.Get()
	defer connPool.Close()
	defer conn.Close()

	reply, errDo := redis.String(conn.Do("get", "address"))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(reply)
}
