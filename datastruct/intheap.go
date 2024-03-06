package datastruct

import (
	"container/heap"
	"fmt"
	"reflect"
)

type IntHeap []int

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {

	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]
	fmt.Println("typeof(old):", reflect.TypeOf(old))
	fmt.Printf("typeof(*h)%T\n:", *h)
	return x
}

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	h := &IntHeap{2, 1, 9, 7, 8, 6}
	heap.Init(h)
	heap.Push(h, 5)
	fmt.Println(h)
	fmt.Println(heap.Pop(h))
}
