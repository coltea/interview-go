// Longest Increasing Subsequence, LIS
package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化 dp 数组
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	fmt.Println(dp)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 1, 2, 5, 3, 7, 101, 18}
	fmt.Println("长度为:", lengthOfLIS(nums)) // 输出: 4
}
