/*
*
  - Definition for a binary tree node.
*/
package leetcode

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	fatherX, fatherY := 0, -1
	levelX, levelY := -1, -2
	lv := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		last := queue[size-1]
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
				if node.Left.Val == y {
					fatherY = node.Val
					levelY = lv
				}
				if node.Left.Val == x {
					fatherX = node.Val
					levelX = lv
				}
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				if node.Right.Val == y {
					fatherY = node.Val
					levelY = lv
				}
				if node.Right.Val == x {
					fatherX = node.Val
					levelX = lv
				}
			}

			if node == last {
				lv++
			}
		}
		if fatherX != 0 && fatherY != -1 {
			break
		}
	}

	fmt.Println(fatherX, fatherY, levelX, levelY)
	if fatherX != fatherY && levelX == levelY {
		return true
	}
	return false

}
func Test1(t *testing.T) {

}
