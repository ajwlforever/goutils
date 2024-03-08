package sort

import (
	"math/rand"
	"time"
)

// genNums1000 生成长度1000的随机数组
func genNums1000() []int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	var nums [1000]int               // Create an array to hold the numbers
	for i := range nums {
		nums[i] = rand.Intn(1000) // Generate a random number from 0 to 999
	}
	return nums[:]
}

// genNums10w 生成长度10万的随机数组
func genNums10w() (nums []int) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	nums = make([]int, 100000)       // Create an array to hold the numbers
	for i := range nums {
		nums[i] = rand.Intn(100000) // Generate a random number from 0 to 100000
	}

	return
}

// genNums1e9 生成长度1亿的随机数组
func genNums1e9() (nums []int) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	nums = make([]int, 1e9)          // Create an array to hold the numbers
	for i := range nums {
		nums[i] = rand.Intn(1e9) // Generate a random number from 0 to 1e9
	}

	return
}
