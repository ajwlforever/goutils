package ratelimit

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func doA(limit Limiter) {
	if !limit.TryAcquire().Ok {
		fmt.Println("reject")
	}
	fmt.Println("do")
}

func TestFixedWindow1(t *testing.T) {
	interval := time.Millisecond * 100 // 0.1s
	ticker := time.NewTicker(interval)
	// 1s 5个请求
	limiter := NewFixedWindowLimiter(time.Second, 5)
	cnt := 0
	for range ticker.C {
		doA(limiter)
		cnt += 1
		if cnt == 1000 {
			ticker.Stop()
			break
		}
	}

}

func TestFixedWindow2(t *testing.T) {
	http.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<h1>hello, world</h1>")
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
	for {
		time.Sleep(time.Second)
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	time.Sleep(time.Second * 5)
	after := time.Now()

	fmt.Println(now)
	fmt.Println(after.Sub(now))
	// sub := after.Sub(now).Seconds()

	duration := time.Duration(after.Sub(now).Seconds() * float64(time.Second))
	fmt.Println("duration:", duration)
	// 判断两个Duration 是否相等，

}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second)
	for {
		<-timer.C
		fmt.Println("timer")
		timer.Reset(time.Second)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Println("timer")
		ticker.Reset(time.Second)
	}
}
