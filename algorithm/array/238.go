package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums)) // [24,12,8,6]

	// nums = []int{-1, 1, 0, -3, 3} [1, -1, -1, 0, 0]
	// fmt.Println(productExceptSelf(nums)) // [0,0,9,0,0]
}

// productExceptSelf 计算除自身以外的所有元素乘积
// 使用前缀积和后缀积的方法，空间复杂度 O(1)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	// 计算左侧所有数字的乘积
	output[0] = 1
	for i := 1; i < n; i++ {
		output[i] = nums[i-1] * output[i-1]
	}

	// 计算右侧所有数字的乘积并更新结果
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}
