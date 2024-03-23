package ratelimit

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RateLimiter struct {
	Pool     chan string
	Cnt      int64 // 处理过的请求总数
	RqCount  int64 // 进来的请求总数
	Cap      int64
	OverTime time.Duration
	locker   sync.Mutex
}

func InitRateLimiter(cap int64, d time.Duration) {
	RL = &RateLimiter{
		Pool:     make(chan string, cap),
		Cnt:      0,
		OverTime: d,
		Cap:      cap,
	}
	var i int64
	for i = 0; i < cap; i++ {
		RL.Pool <- "init" // 初始化限流池，
	}
}

var RL *RateLimiter

func doSomething() {
	reqCountAdd()
	select {
	case omit := <-RL.Pool:
		_ = omit
		fmt.Println("doChat", RL.Cnt)
		go execSomething() //异步执行
	case <-time.After(RL.OverTime):
		fmt.Println("reject ")
	}
}

func reqCountAdd() {
	defer RL.locker.Unlock()
	RL.locker.Lock()
	RL.RqCount += 1
	fmt.Println("reqCount:", RL.RqCount)
}

func execSomething() {
	defer func() {
		RL.locker.Lock()
		RL.Cnt += 1
		RL.locker.Unlock()
		RL.Pool <- "omit"
	}()
	seed := rand.Float64() * 2
	t := seed * (float64(RL.OverTime.Nanoseconds())) // 模拟 时间在超时左右
	time.Sleep(time.Duration(t))
}
