package main

import (
	"fmt"
)

// 归并排序函数
func mergeSort(arr []int) []int {
	// 如果数组长度小于等于1，不需要排序，直接返回
	if len(arr) <= 1 {
		return arr
	}

	// 分割数组
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // 递归排序左半部分
	right := mergeSort(arr[mid:]) // 递归排序右半部分

	// 合并左右两个已排序的部分
	return merge(left, right)
}

// 合并两个已排序的数组
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	// 合并两个数组
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果左半部分还有剩余
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// 如果右半部分还有剩余
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	// 测试归并排序
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("原始数组:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("排序后的数组:", sortedArr)
}
