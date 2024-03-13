package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type Score struct {
	diff  int
	index int // 存取maxDiff时，还存储原来所在的index，好帮助累积分值
}
type ScoreSlice []Score

func (a ScoreSlice) Len() int           { return len(a) }
func (a ScoreSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ScoreSlice) Less(i, j int) bool { return a[i].diff < a[j].diff }

func stoneGameVI1(aliceValues []int, bobValues []int) int {
	alice, bob := 0, 0
	diff := make(ScoreSlice, len(aliceValues))
	for i := 0; i < len(aliceValues); i++ {
		diff[i].diff = aliceValues[i] + bobValues[i]
		diff[i].index = i
	}
	sort.Sort(diff)

	last := len(aliceValues) - 1
	turn := 0
	for {
		if last < 0 {
			break
		}
		if turn == 0 {
			//alice
			// 选diff最大的
			index := diff[last].index
			alice += aliceValues[index]
			last--
			turn = 1

		} else {
			// bob
			index := diff[last].index
			bob += bobValues[index]
			last--
			turn = 0
		}
	}
	fmt.Printf("alice:%d bob:%d \n", alice, bob)
	if alice > bob {
		return 1
	} else if alice < bob {
		return -1
	} else {
		return 0
	}

}

func TestMain(m *testing.M) {
	fmt.Println(stoneGameVI1([]int{3, 1}, []int{2, 1}))
	fmt.Println(stoneGameVI1([]int{1, 2}, []int{3, 1}))
	fmt.Println(stoneGameVI1([]int{2, 4, 3}, []int{1, 6, 7}))
}
