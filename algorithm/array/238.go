package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums)) // [24,12,8,6]

	// nums = []int{-1, 1, 0, -3, 3} [1, -1, -1, 0, 0]
	// fmt.Println(productExceptSelf(nums)) // [0,0,9,0,0]
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	output[0] = 1
	for i := 1; i < n; i++ {
		output[i] = nums[i-1] * output[i-1]
	}
	fmt.Println(output)
	right := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}
