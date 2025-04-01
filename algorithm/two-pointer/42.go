package main

import (
	"fmt"
)

func main() {

	// fmt.Println(trap([]int{3, 4, 2, 6, 3, 5, 2})) //                   // 9
	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	// fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))             // 9
	// fmt.Println(trap([]int{5, 4, 1, 2}))                   // 1
	// fmt.Println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3})) // 23
	fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3})) // 3
}

// trap 计算能接住的雨水总量
// 使用双指针法，时间复杂度 O(n)，空间复杂度 O(1)
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	totalWater := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}

	return totalWater
}
