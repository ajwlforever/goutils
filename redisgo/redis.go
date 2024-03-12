package redisgo

import (
	"context"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	// create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
