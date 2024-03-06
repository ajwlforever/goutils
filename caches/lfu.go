package caches

import (
	"container/heap"
	"fmt"
	"sync"
)

// LFU

// 优先队列
type Item struct {
	cnt   int // 访问次数
	value any
	key   any
	index int // 记录在优先队列的位置
}

type PriorityQueue []*Item                 // 优先队列，采用小顶堆，访问次数少的位于堆顶
func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(a, b int) bool { return q[a].cnt < q[b].cnt }
func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index, q[j].index = q[j].index, q[i].index
}

func (q *PriorityQueue) Pop() any {
	n := len(*q)
	old := *q
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q *PriorityQueue) Push(item any) {
	n := len(*q)
	x := item.(*Item)
	x.index = n
	*q = append(*q, x)
}

// 更新访问次数
func (q *PriorityQueue) Update(item *Item, c int) {
	item.cnt = c
	heap.Fix(q, item.index)
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{
		cap:   cap,
		cache: make(map[any]*Item, 0),
		queue: &PriorityQueue{},
	}
}

type LFUCache struct {
	queue  *PriorityQueue
	cache  map[any]*Item
	cap    int
	rwLock sync.RWMutex
}

// Get exists return val else return nil
func (lfu *LFUCache) Get(key any) (val any) {
	lfu.rwLock.RLock()
	defer lfu.rwLock.RUnlock()
	node := lfu.cache[key]
	if node == nil {
		fmt.Printf("no k-v :%v-%v\n", key, val)
		return
	}
	val = node.value
	lfu.queue.Update(node, node.cnt+1)
	fmt.Printf("get k-v :%v-%v\n", key, val)
	return

}

// Push  to cache
//
//	 if exists update k-v
//		else if full delete lfu
//		not full->   insert
func (lfu *LFUCache) Push(key any, val any) {
	// 写锁
	lfu.rwLock.Lock()
	defer lfu.rwLock.Unlock()

	pool := lfu.cache
	queue := lfu.queue
	if pool[key] != nil {
		// update cache
		node := pool[key]
		node.value = val
		fmt.Printf("update k-v:%v-%v\n", key, val)
	} else {
		if len(pool) >= lfu.cap {
			// 删除最不经常用的缓存
			if err := lfu.deleteMinCnt(); err != nil {
				return
			} //
		}
		item := &Item{
			cnt:   0,
			key:   key,
			value: val,
		}
		// 放进缓存
		heap.Push(queue, item)
		pool[key] = item
		fmt.Printf("push k-v :%v-%v\n", key, val)
	}

}

func (lfu *LFUCache) deleteMinCnt() error {

	q := *lfu.queue
	delete(lfu.cache, q[0].key)
	*lfu.queue = q[1:len(q)]

	return nil
}

func (lfu LFUCache) lazyDelete() error {
	// TODO
	return nil
}

func (lfu LFUCache) delete(item *Item) error {
	// TODO
	return nil
}
