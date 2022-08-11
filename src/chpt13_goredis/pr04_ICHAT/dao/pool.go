package dao

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var ConnPool *redis.Pool

func init() {
	ConnPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
