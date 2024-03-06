package datastruct

import (
	list2 "container/list"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	list := NewLinkedList()
	list.InsertToBacks(1, 3, 4, 56, 7, 9, 6, 3)
	list.print()
}

func TestRemoveNode(t *testing.T) {
	list := NewLinkedList()
	list.InsertToBacks(1, 3, 4, 56, 7, 9, 6, 3)

	list.print()
}

func TestContainer(t *testing.T) {
	l := list2.New()
	l.PushFront("s")
	l.PushFront(1)
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		fmt.Printf("%T\n", e.Value)
	}
}
