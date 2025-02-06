// Longest Increasing Subsequence, LIS
package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	for i, num := range nums {

	}
	return maxLen
}

func main() {
	nums := []int{10, 1, 2, 5, 3, 7, 101, 18}
	fmt.Println("长度为:", lengthOfLIS(nums)) // 输出: 4
}
