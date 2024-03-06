package datastruct

import (
	list2 "container/list"
	"fmt"
)

func testList() {
	arrs := []int{
		1, 2, 3, 4, 5, 6,
	}
	list := list2.New()
	for _, v := range arrs {
		list.PushFront(v)
	}

	printList(list)

}

func printList(list *list2.List) {
	l := list
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}
