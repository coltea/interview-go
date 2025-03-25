package main

import (
	"fmt"
)

//
// N = 3, W = 4
// weight = [2, 1, 3]
// value = [4, 2, 3]

func knapsack(N, W int, weight, value []int) int {
	dp := make([]int, W+1)

	for i := 1; i <= N; i++ {
		fmt.Println(i)
		for w := W; w >= weight[i-1]; w-- { // 逆序遍历，防止覆盖
			fmt.Println(dp[w-weight[i-1]], " ", value[i-1])
			dp[w] = max(dp[w], dp[w-weight[i-1]]+value[i-1])
			fmt.Println(dp)
		}
	}

	return dp[W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{2, 1, 3}
	value := []int{4, 2, 3}
	fmt.Println(knapsack(3, 4, weight, value)) // 输出 6
}
