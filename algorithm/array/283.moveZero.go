package main

import (
	"fmt"
)

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	i := 0
	i2 := 0

	for i2 < len(nums) {
		if nums[i2] != 0 {
			nums[i], nums[i2] = nums[i2], nums[i]
			fmt.Println(nums)
			i++
		}
		i2++
	}
}
