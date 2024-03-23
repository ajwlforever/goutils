package datastruct

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 9, 7, 8, 6}
	heap.Init(h)
	heap.Push(h, 5)
	fmt.Println(h)
	fmt.Println(heap.Pop(h))
}

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	} // Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	pp(pq)

	item := &Item{
		value:    "orange",
		priority: 1,
		index:    -1,
	}

	heap.Push(&pq, item)
	pp(pq)
	pq.update(item, "pipe", 10)
	pp(pq)
}

func TestSort1(t *testing.T){
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
    sort.Sort(sort.Reverse(sort.IntSlice(s)))
    fmt.Println(s)
}