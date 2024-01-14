package main

import (
	"context"
	"fmt"
	"time"
)

func do_thing(ctx context.Context) {

	ticker := time.NewTicker(time.Second) //每1s触发一次ticker
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("over 0")
			return
		default:
			fmt.Println("do")
		}
	}
}

func do_thing2(ctx context.Context) {
	ticker := time.NewTicker(time.Second) //每1s触发一次ticker
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("over2")
			return
		default:
			fmt.Println("do2")
		}
	}
}
func CancelHandle() {
	ctx, cancel := context.WithCancel(context.Background())

	//现在这个ctx，是一个WithCancel的Context，cancel触发即结束所有子context
	go do_thing(ctx) // 子节点以ctx为上下文启动
	time.Sleep(10 * time.Second)
	cancel()
}
func DeadlineHandle() {
	parent := context.Background()
	five := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(parent, five)

	go do_thing(ctx)
	go do_thing2(parent)
	time.Sleep(10 * time.Second)
	cancel() // 10s父节点结束，但5s子节点应该位于deadline
}

func TimeoutHandle() {
	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)

	go do_thing(ctx)
	// go do_thing2(parent)
	time.Sleep(10 * time.Second)
	cancel() // 10s父节点结束，但5s子节点应该位于deadline
}
func main() {
	// CancelHandle()
	//DeadlineHandle()
	TimeoutHandle()
}
