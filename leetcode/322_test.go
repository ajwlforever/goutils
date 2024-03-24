package leetcode

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestCoin(t *testing.T) {
	// coins := []int{1, 2, 5}
	// amount := 11
	// fmt.Println(coinChange(coins, amount))

	coins := []int{186, 419, 83, 408}
	amount := 6249
	fmt.Println(coinChange1(coins, amount))
}
func coinChange(coins []int, amount int) (res int) {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // 降序
	for _, v := range coins {

		res += amount / v
		amount = amount % v // 剩余的钱
		if amount == 0 {
			return res
		}
	}
	if amount != 0 {
		return -1
	}
	return res
}

func coinChange1(coins []int, amount int) (res int) {
	memo := make([]int, amount+1) // 備忘錄記錄
	for i := 0; i <= amount; i++ {
		memo[i] = -2
	}
	return dp(coins, amount, memo)
}

func dp(coins []int, amount int, memo []int) (res int) {
	//boundary check
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// memo check
	if memo[amount] != -2 {
		return memo[amount]
	}
	res = math.MaxInt32
	// memo
	for _, coin := range coins {
		// select this coin
		tmp := dp(coins, amount-coin, memo)
		if tmp == -1 {
			continue
		}
		res = min(tmp+1, res)
	}
	// 把计算结果存入备忘录
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return

}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
