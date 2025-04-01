package main

import (
	"fmt"
)

func main() {

	// fmt.Println(trap([]int{3, 4, 2, 6, 3, 5, 2})) //                   // 9
	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))             // 9
	fmt.Println(trap([]int{5, 4, 1, 2}))                   // 1
	fmt.Println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3})) // 23
	fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))          // 3
}

func trap(height []int) int {
	total := 0
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				// fmt.Println(min(leftMax, height[right]) - height[left])
				// fmt.Println(leftMax, height[right], height[left])
				total += min(leftMax, height[right]) - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				total += min(rightMax, height[left]) - height[right]
			}
			right--
		}
	}
	return total
}
