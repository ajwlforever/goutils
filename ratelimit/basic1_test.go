package ratelimit

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	NewRL(2) // 初始化大小为2的限流器
	for {
		doChat() // 开始模拟一直发送请求
	}
}

func Test2(t *testing.T) {
	InitRateLimiter(2, time.Second*1)

	ticker := time.NewTicker(300 * time.Millisecond)
	for range ticker.C {
		go doSomething()
	}
	time.Sleep(time.Hour)
}
