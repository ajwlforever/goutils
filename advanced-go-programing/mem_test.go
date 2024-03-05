package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestBingfa1(t *testing.T) {
	done := make(chan int)

	go func() {
		fmt.Println("你好, 世界")
		<-done
	}()

	done <- 1
}

func TestCacheChannel(t *testing.T) {
	done := make(chan int, 10) // 带 10 个缓存

	// 开 N 个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("你好, 世界" + strconv.Itoa(i))
			done <- 1
		}(i)
	}

	// 等待 N 个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	cnt := 10
	wg.Add(cnt)

	for i := 0; i < cnt; i++ {
		go func(i int) {
			fmt.Println("你好, 世界" + strconv.Itoa(i))
			wg.Done()
		}(i)
	}

	wg.Wait()
}

type ErrorGroup struct {
}
