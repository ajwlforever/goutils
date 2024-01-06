package syncs

import (
	"sync"
	"testing"
	"time"
)

// 执行一万次结果都不会改变
func TestMutex(t *testing.T) {
	var count *int64
	mutex := &sync.Mutex{}
	count = new(int64)
	*count = 1
	go Mutex1(mutex, count)
	go Mutex2(mutex, count)
	time.Sleep(10000)
}

func TestNoMutex(t *testing.T) {
	var count int64 = 1
	go NoMutex1(&count)
	go NoMutex2(&count)
	time.Sleep(10000)
}
