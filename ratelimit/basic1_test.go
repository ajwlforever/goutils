package ratelimit

import (
	"fmt"
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

func TestGenerateToken(t *testing.T) {
	count := 1000000 //
	set := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		str := generateToken()
		// fmt.Println(str)
		if set[str] {
			fmt.Println("我擦你", str)
		} else {
			set[str] = true
		}
	}
}

func BenchMarkGenerateToken(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = i
		count := 1000000 //
		set := make(map[string]bool, count)
		for i := 0; i < count; i++ {
			str := generateToken()
			// fmt.Println(str)
			if set[str] {
				fmt.Println("我擦你", str)
			} else {
				set[str] = true
			}
		}
	}
	b.StopTimer()

}
