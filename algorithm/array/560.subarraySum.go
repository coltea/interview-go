package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))  // 2
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))  // 2
	fmt.Println(subarraySum([]int{1, -1, 0}, 0)) // 3
}

func subarraySum(nums []int, k int) int {
	res := 0
	for i := range nums {
		i2 := i

		sum := 0
		for i2 >= 0 {
			sum += nums[i2]
			if sum == k {
				res++
			}
			i2--
		}
	}

	return res
}
