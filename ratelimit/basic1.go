package ratelimit

import (
	"fmt"
	"time"
)

var rlChannel chan int
var cnt int64

// doChat 模拟AI服务接口
func doChat() {
	rlChannel <- 1 // 请求进入限流器缓冲池
	fmt.Println("doChat", cnt)
	go execChat() //异步执行
}

func execChat() {
	defer func() {
		<-rlChannel // 释放限流器的令牌
		cnt += 1
	}()
	time.Sleep(time.Second * 5) // 1s模拟接口时间
}

func NewRL(cap int) {
	rlChannel = make(chan int, cap)
	cnt = 0
}
