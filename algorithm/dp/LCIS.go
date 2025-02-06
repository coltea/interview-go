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

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("最长连续递增子序列长度:", findLengthOfLCIS(nums)) // 输出: 3

	nums = []int{2, 2, 2, 2, 2}
	fmt.Println("最长连续递增子序列长度:", findLengthOfLCIS(nums)) // 输出: 1
}
