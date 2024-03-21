package redisgo

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-redis/redis/v8"
)

func Script1() {
	rdb := NewRedisClient()
	f, _ := os.Open("tokenbucket.lua")
	s, _ := io.ReadAll(f)

	IncrByXX := redis.NewScript(string(s))
	ctx := context.Background()
	n, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)
}
