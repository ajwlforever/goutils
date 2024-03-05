package main

import (
	"fmt"
)

// LRU 最近最少使用Cache “如果数据最近被访问过，那么将来被访问的几率也更高“

// Item 最近最少使用的

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

type LinkedList struct {
	len  int
	root *Node
}
type LRUCache struct {
	cap   int
	list  *LinkedList   // doubleLinkedList
	cache map[int]*Node // k-v
}

func (list LinkedList) DeleteNode(node *Node) {
	list.RemoveNode(node)
	node.next = nil // for save
	node.prev = nil
}

func (list LinkedList) RemoveNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (list LinkedList) MoveToFirst(node *Node) {
	list.RemoveNode(node)
	list.PutToFirst(node)
}
func (list LinkedList) PutToFirst(node *Node) {
	root := list.root
	nxt := root.next
	nxt.prev = node
	root.next = node
	node.prev = root
	node.next = nxt
}
func (lru LRUCache) Put(k, v int) {

	node := lru.cache[k]
	if node != nil {
		// exists, update k-v
		node.val = v
		node.key = k
		// move node to first
		lru.list.MoveToFirst(node)

	} else {
		// dont exists, cap
		if lru.cap <= lru.list.len {
			// cache full, delete the recently use min  --> queue last
			root := lru.list.root
			last := root.prev
			// delete last in list
			lru.list.DeleteNode(last)

			// delete last in map
			delete(lru.cache, last.key)
			lru.list.len = lru.list.len - 1
		}
		// put to the first
		node = &Node{
			key: k,
			val: v,
		}
		lru.list.PutToFirst(node)
		lru.list.len = lru.list.len + 1
		lru.cache[k] = node

	}

}

func (lru LRUCache) Get(k int) (v int) {
	node := lru.cache[k]
	if node != nil {
		fmt.Println("k-v:", k, " ", node.val)
		return node.val
	}
	fmt.Println("不存在k:", k)
	return -1
}

func New(cap int) LRUCache {
	node := &Node{
		key: -1,
		val: -1,
	}
	node.prev = node
	node.next = node
	return LRUCache{
		cache: map[int]*Node{},
		cap:   cap,
		list: &LinkedList{
			len:  0,
			root: node,
		},
	}
}

func main() {
	c := New(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
	c.Put(3, 3)
	c.Get(1)
	c.Put(1, 1)
	c.Put(4, 4)
	c.Get(3)
}
