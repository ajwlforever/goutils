package datastruct

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

func NewLinkedList() *LinkedList {
	node := &Node{
		key: -1,
		val: -1,
	}
	node.prev = node
	node.next = node
	return &LinkedList{
		len:  0,
		root: node,
	}
}

func (list LinkedList) Len() int {
	return list.len
}

func (list LinkedList) InsertToBacks(vals ...int) {
	for _, val := range vals {
		list.InsertToBack(val)
	}
}

// InsertToBack 插入尾部
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
	list.len = list.len + 1
}

// InsertNodeToBack
func (list LinkedList) InsertNodeToBack(node *Node) {
	r := list.root
	l := r.prev
	l.next = node
	r.prev = node
	node.next = r
	node.prev = l
	list.len = list.len + 1
}

// InsertAt 插入到指定位置
func (list LinkedList) InsertAt(val int, loc int) {

}
func (list LinkedList) InsertAfter(v int, node *Node) {

}
func (list LinkedList) Get(loc int) *Node {
	if loc >= list.len {
		return nil
	}
	l := loc
	node := list.root
	l -= 1
	for l > 0 {
		l -= 1
		node = node.next
	}
	return node
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
	fmt.Printf("DoubleLinkedList:  ")
	root := list.root
	if root.prev != nil {
		fmt.Printf("< ")
	}
	fmt.Printf("%v ", root.val)
	if root.next != nil {
		fmt.Printf(" >")
	}
	tmp := root.next
	for tmp != root {
		if tmp.prev != nil {
			fmt.Printf("< ")
		}
		fmt.Printf("%v ", tmp.val)
		if tmp.next != nil {
			fmt.Printf(" >")
		}
		tmp = tmp.next
	}

}

//
//func main() {
//	list := New()
//	list.InsertToBacks(1, 3, 4, 5, 9, 8, 7)
//	list.print()
//}
