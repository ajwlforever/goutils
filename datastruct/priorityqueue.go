package datastruct

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // 实际值
	priority int    // 优先级
	index    int    // 在堆中的位置
}

// PriorityQueue 使用heap实现优先队列
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push heap implementation
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop  heap implementation
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	old := *pq
	x := old[n-1]
	// avoid mem leak
	old[n-1] = nil
	x.index = -1 // for safe
	*pq = old[:n-1]
	return x
}

// update update the item of PriorityQueue
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

//
//func main() {
//	// Some items and their priorities.
//	items := map[string]int{
//		"banana": 3, "apple": 2, "pear": 4,
//	} // Create a priority queue, put the items in it, and
//	// establish the priority queue (heap) invariants.
//	pq := make(PriorityQueue, len(items))
//	i := 0
//	for value, priority := range items {
//		pq[i] = &Item{
//			value:    value,
//			priority: priority,
//			index:    i,
//		}
//		i++
//	}
//
//	heap.Init(&pq)
//	pp(pq)
//
//	item := &Item{
//		value:    "orange",
//		priority: 1,
//		index:    -1,
//	}
//
//	heap.Push(&pq, item)
//	pp(pq)
//	pq.update(item, "pipe", 10)
//	pp(pq)
//}

// pp print PriorityQueue
func pp(pq PriorityQueue) {
	for _, v := range pq {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
