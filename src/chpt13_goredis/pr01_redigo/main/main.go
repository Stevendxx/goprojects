package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// redis

/*	1. key
	DEL
	DUMP
	EXISTS
	EXPIRE
	EXPIREAT
	KEYS
	MIGRATE
	MOVE
	OBJECT
	PERSIST
	PEXPIRE
	PEXPIREAT
	PTTL
	RANDOMKEY
	RENAME
	RENAMENX
	RESTORE
	SORT
	TTL
	TYPE
	SCAN
*/

/*	2. string
	APPEND
	BITCOUNT
	BITOP
	DECR
	DECRBY
	GET
	GETBIT
	GETRANGE
	GETSET
	INCR
	INCRBY
	INCRBYFLOAT
	MGET
	MSET
	MSETNX
	PSETEX
	SET
	SETBIT
	SETEX
	SETNX
	SETRANGE
	STRLEN
*/

/*	3. hash
	HDEL
	HEXISTS
	HGET
	HGETALL
	HINCRBY
	HINCRBYFLOAT
	HKEYS
	HLEN
	HMGET
	HMSET
	HSET
	HSETNX
	HVALS
	HSCAN
*/

/*	4. list
	BLPOP
	BRPOP
	BRPOPLPUSH
	LINDEX
	LINSERT
	LLEN
	LPOP
	LPUSH
	LPUSHX
	LRANGE
	LREM
	LSET
	LTRIM
	RPOP
	RPOPLPUSH
	RPUSH
	RPUSHX
*/

/*	5. set
	SADD
	SCARD
	SDIFF
	SDIFFSTORE
	SINTER
	SINTERSTORE
	SISMEMBER
	SMEMBERS
	SMOVE
	SPOP
	SRANDMEMBER
	SREM
	SUNION
	SUNIONSTORE
	SSCAN
*/

/*	6. sorted set
	ZADD
	ZCARD
	ZCOUNT
	ZINCRBY
	ZRANGE
	ZRANGEBYSCORE
	ZRANK
	ZREM
	ZREMRANGEBYRANK
	ZREMRANGEBYSCORE
	ZREVRANGE
	ZREVRANGEBYSCORE
	ZREVRANK
	ZSCORE
	ZUNIONSTORE
	ZINTERSTORE
	ZSCAN
*/

// redis key
func KeyTest(conn redis.Conn) {
	// get
	reply, errDo := redis.String(conn.Do("get", "address"))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(reply)
}

// redis hash
func HashTest(conn redis.Conn) {
	// hgetall
	user01, errDo := redis.StringMap(conn.Do("hgetall", "user01"))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(user01)
}

// redis list
func ListTest(conn redis.Conn) {
	// lrange
	list01, errDo := redis.Strings(conn.Do("lrange", "list01", 0, -1))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(list01)
}

// redis set
func SetTest(conn redis.Conn) {
	// smembers
	set01, errDo := redis.Ints(conn.Do("smembers", "set01"))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(set01)
}

// redis sorted set
func SsetTest(conn redis.Conn) {
	// zrange
	sset01, errDo := redis.StringMap(conn.Do("zrange", "sset01", 0, -1, "withscores"))
	if errDo != nil {
		fmt.Println("操作失败！")
	}
	fmt.Println(sset01)
}

func main() {
	conn, errDial := redis.Dial("tcp", "localhost:6379")
	if errDial != nil {
		panic(errDial)
	}
	defer conn.Close()

	// 1. key
	//KeyTest(conn)

	// 2. hash
	//HashTest(conn)

	// 3. list
	//ListTest(conn)

	// 4. set
	//SetTest(conn)

	// 4. sorted set
	SsetTest(conn)
}
