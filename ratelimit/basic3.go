package ratelimit

import (
	"math/rand"
	"sync"
	"time"
)

type Limiter1 interface {
	tryAcquire() bool
	tryRelease() bool
	acquire() (token string)
	release(token string)
}

// 加入令牌桶算法的限流器

// TokenRateLimiter1
// Cap 是桶的容量
// Rate 规定处理请求的速度，以interval为单位
// ReqRate 实际请求进来的速率，以interval为单位，Rate > ReqRate 时放行。
// Pool 令牌池，存放令牌 --》
// Interval 检测令牌桶是否正常的区间
type TokenRateLimiter1 struct {
	Cap      int64
	Rate     int64
	ReqRate  int64           // Rate > ReqRate 时放行。
	Pool     map[string]bool // 令牌池
	Lock     sync.Mutex
	Interval time.Duration //
	timer    *time.Timer   //定时器
	log      Logger
}

// Logger 限流器的日志记录
type Logger struct {
}

func NewTokenRateLimiter1(cap int64, rate int64, interval time.Duration) *TokenRateLimiter1 {

	t := &TokenRateLimiter1{
		Cap:      cap,
		Rate:     rate,
		Interval: interval,
		Pool:     make(map[string]bool),
		timer:    time.NewTimer(interval),
		log:      Logger{},
	}
	// 启用监控模式
	go t.watchDog()
	return t
}

func (limiter *TokenRateLimiter1) watchDog() {
	// TODO
	for {
		select {
		case <-limiter.timer.C:
			// Reste something
			// limiter.timer.Reset(limiter.Interval)
			limiter.ReqRate = 0
		}
	}

}

// tryAcuire 尝试是否可以获取令牌，返回true or false
func (limiter *TokenRateLimiter1) tryAcquire() bool {
	return limiter.ReqRate < limiter.Rate
}

// 随机8位字符串
func generateToken() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = byte(rand.Int31n(150))
	}
	return string(b)
}

// acquire 会请求一个令牌， token=="-1"说明请求失败
// 1. 检查是否放行，给予令牌： 放行的条件： rate =
func (limiter *TokenRateLimiter1) acquire() (token string) {
	limiter.Lock.Lock()
	defer limiter.Lock.Unlock()
	if limiter.tryAcquire() {
		// 给予令牌 
		return 
	}
	if 
	// 

	return 
}
