package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b) // a
		if n == 0 {
			break
		} else {
			m := make(map[int]int, 0)
			all := make(map[int]int, 0)
			res := 0
			reader := bufio.NewReader(os.Stdin)
			reader.ReadLine()
			for i := 0; i < b; i++ {
				data, _, _ := reader.ReadLine()
				str := string(data)
				s := strings.Split(str, " ")
				t1, _ := strconv.Atoi(s[0])
				t2, _ := strconv.Atoi(s[1])
				t3 := s[2]
				all[t1] = 1
				all[t2] = 1
				if t3 == "R" {
					m[t1] = 1
					m[t2] = 1
				}
			}
			for i := 1; i <= a; i++ {
				if all[i] == m[i] {
					res++
				}
			}
			fmt.Println(res)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4
// 3 4 1 2
func canSorted(lists []*ListNode) (res []bool) {
	// write code here
	n := len(lists)
	res = make([]bool, n)
	for i := 0; i < n; i++ {
		l1, r1, l2, r2 := 0, 0, 0, 0
		list := lists[i]
		l1 = list.Val
		list = list.Next
		pre := l1
		for list != nil {
			if pre < list.Val {
				pre = list.Val
			} else {
				l2 = list.Val
				pre = list.Val
				list = list.Next
				break
			}
			list = list.Next
		}
		_ = r1
		for list != nil {
			if pre < list.Val {
				pre = list.Val
			} else {
				pre = list.Val
				break
			}
		}
		if list != nil{
			res[i] = false
			return 
		}else{
			r2 = list.Val
		}
		if r1 == 0 && r2 == 0 && l2 == 0 {
			res[i] = true
			continue
		}
		if r2 < l1 {
			res[i] = true
			continue
		}
		res[i] = false
	}

	return
}


func canSorted(lists []*ListNode) (res []bool) {
	// write code here
	n := len(lists)
	res = make([]bool, n)
	for i := 0; i < n; i++ {
		l1, r1, l2, r2 := 0, 0, 0, 0
		list := lists[i]
		l1 = list.Val
		list = list.Next
		pre := l1
		for list != nil {
			if pre < list.Val {
				pre = list.Val
			} else {

				l2 = list.Val
				pre = list.Val
				list = list.Next
				break
			}
			list = list.Next
		}
		_ = r1
		for list != nil {
			if pre < list.Val {
				pre = list.Val
			} else {
				pre = list.Val 
				break
			}
		}
        if list != nil{
			res[i] = false
			return 
		}else{
			r2 = pre
		}
        //fmt.Println(l1,r1,l2,r2)
		if (r1 == 0 && r2 == 0 && l2 == 0) || (r1 == 0 && l2 == 0) {
			res[i] = true
			continue
		}
		if r2 < l1 {
			res[i] = true
			continue
		}
		res[i] = false;
	}

	return
}