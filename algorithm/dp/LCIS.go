package main

import (
	"fmt"
)

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLen := 1
	currLen := 1

	// 遍历数组
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			currLen = 1
		}
	}

	return maxLen
}

func lengthOfLIS2(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
		fmt.Println(dp)

	}

	return maxLength
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(lengthOfLIS2(nums))

	nums = []int{2, 2, 2, 2, 2}
	fmt.Println(lengthOfLIS2(nums))

	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS2(nums))

}
