package config

import "github.com/redis/go-redis/v9"

func NewRedis() redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return rdb
}
