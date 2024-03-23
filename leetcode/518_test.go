package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func Test518(t *testing.T) {
	// dp := make([]int, 10)
	// for i := range dp {
	// 	fmt.Printf("%v\n", i)
	// }
	// for i, num := range dp {
	// 	fmt.Printf("%v %v\n ", i, num)
	// }
	amount := 5
	conins := []int{1, 2, 5}
	fmt.Println(change(amount, conins))
}

func change(amount int, coins []int) int {
	sort.Ints(coins)
	n := len(coins)
	dp := make([][]int, n+1) // dp[n][amount] 表示前n个钱币凑amount的方案数
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n-1][n-1]
}
