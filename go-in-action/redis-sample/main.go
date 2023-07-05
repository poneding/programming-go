package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisPool = newRedisPool()

func main() {
	client := redisPool.Get()
	defer client.Close()

	_, e := client.Do("SET", "redigo", "redigo value")
	if e != nil {
		panic(e)
	}

	val, e := client.Do("GET", "redigo")
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s\n", val)

	client.Do("ZADD", "redigo-zset", 4, "car")
	client.Do("ZADD", "redigo-zset", 2, "bike")

	zsetVal, _ := client.Do("ZRANGE", "redigo-zset", 0, -1, "WITHSCORES")
	fmt.Printf("%s,\n", zsetVal)
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   60,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial("tcp", "172.30.37.246:6379")
			if e != nil {
				panic("dial error:" + e.Error())
			}
			return c, e
		},
	}
}
