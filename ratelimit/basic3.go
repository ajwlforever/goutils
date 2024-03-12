package ratelimit

import (
	"sync"
	"time"
)

// 加入令牌桶算法的限流器

// TokenRateLimiter
// Cap 是桶的容量
// Rate 是处理请求的速度，以s为单位
// Pool
// Interval 检测令牌桶是否正常的区间
// ReqRate 请求进来的数量
type TokenRateLimiter struct {
	Cap      int64
	Rate     int64
	Pool     map[string]bool // 令牌池
	Lock     sync.Mutex
	Interval time.Duration //
	timer    *time.Timer   //定时器
	ReqRate  int64
}

func NewTokenRateLimiter(cap int64, rate int64, interval time.Duration) *TokenRateLimiter {

	t := &TokenRateLimiter{
		Cap:      cap,
		Rate:     rate,
		Interval: interval,
		Pool:     make(map[string]bool),
		timer:    time.NewTimer(interval),
	}
	// 启用监控模式
	go watchDog(t)
	return t
}

func watchDog(limiter *TokenRateLimiter) {

}
