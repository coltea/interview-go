// Longest Increasing Subsequence, LIS
package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

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

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("最长连续递增子序列长度:", lengthOfLIS(nums)) // 输出: 3

	nums = []int{2, 2, 2, 2, 2}
	fmt.Println("最长连续递增子序列长度:", lengthOfLIS(nums)) // 输出: 1

	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("最长连续递增子序列长度:", lengthOfLIS(nums)) // 输出: 1

}
