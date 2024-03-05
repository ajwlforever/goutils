package main

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	i := 1
	for {
		size := len(nums)
		if i == size {
			break
		}
		if nums[i] == nums[i-1] { // 删掉nums[i]
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++ // 没有重复元素才skip
		}
	}
	for _, num := range nums {
		fmt.Print(num)
	}
	fmt.Println()
	return len(nums)
}

func Test26(t *testing.T) {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0}
	fmt.Println(removeDuplicates(nums))

}
