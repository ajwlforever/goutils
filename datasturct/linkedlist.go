package main

import (
	"fmt"
)

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

func (list LinkedList) Len() int {
	return list.len
}

// Insert 插入尾部
func (list LinkedList) InsertToBack(val int) {
	r := list.root
	l := r.prev
	node := &Node{
		val:  val,
		next: r,
		prev: l,
	}
	l.next = node
	r.prev = node
}

// InsertAt 插入到指定位置
func (list LinkedList) InsertAt(val int, loc int) {

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
	list.InsertToFirst(node)
}
func (list LinkedList) InsertToFirst(node *Node) {
	root := list.root
	nxt := root.next
	nxt.prev = node
	root.next = node
	node.prev = root
	node.next = nxt
}

func (list LinkedList) print() {
	if list.root == nil {
		return
	}
	root := list.root
	fmt.Printf("DoubleLinkedList:%v <> ", root.val)
	tmp := root.next
	for tmp != root {
		fmt.Printf("%v <> ", tmp.val)
		tmp = tmp.next
	}

}
func main() {

}
