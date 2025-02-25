package main

import (
	"fmt"
)

func main() {
	k := []int{1, 2, 3, 4, 5, 6}
	rotate(k, 3)
	fmt.Println(k)
}

func rotate(nums []int, k int) {
	index := 0
	for index < len(nums) {
		cursor := index - k
		if cursor < 0 {
			cursor += len(nums)
		}
		nums[index] = nums[cursor]
		index++
	}
}
