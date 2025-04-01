package main

// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

import (
	"fmt"
)

func main() {
	// fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	// fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	// fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5})) // 3
}

func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, len(nums)+1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// 只要当前窗口内的 sum ≥ target，就尝试缩小窗口
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left] // 缩小窗口
			left++
		}
	}

	// 如果没有找到符合条件的子数组，返回 0
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
