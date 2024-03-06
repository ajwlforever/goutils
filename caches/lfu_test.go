package caches

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	cache := NewLFUCache(2)
	cache.Push(1, 1)
	cache.Push(2, 2)
	cache.Get(1)
	cache.Get(1)
	cache.Get(1)
	cache.Push(3, 3)
	cache.Get(2)
}

func TestPq(t *testing.T) {
	item := &Item{
		cnt:   2,
		value: 1,
		key:   1,
		index: 0,
	}
	q := make(PriorityQueue, 0)
	heap.Push(&q, item)
	item2 := &Item{
		cnt:   1,
		value: 1,
		key:   1,
		index: 0,
	}
	heap.Push(&q, item2)
	for _, v := range q {
		fmt.Println(v)
	}
}
