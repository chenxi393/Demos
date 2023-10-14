package example

import (
	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	RedisClient = rdb
}
