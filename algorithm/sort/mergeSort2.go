package main

import (
	"fmt"
)

func mergeSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := mergeSort2(arr[:mid])
	rightArr := mergeSort2(arr[mid:])
	return merge2(leftArr, rightArr)
}

func merge2(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	i1, i2 := 0, 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] < arr2[i2] {
			res = append(res, arr1[i1])
			i1++
		} else {
			res = append(res, arr2[i2])
			i2++
		}
	}

	for i1 < len(arr1) {
		res = append(res, arr1[i1])
		i1++
	}
	for i2 < len(arr2) {
		res = append(res, arr2[i2])
		i2++
	}

	return res
}

func main() {
	// 测试归并排序
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("原始数组:", arr)

	sortedArr := mergeSort2(arr)
	fmt.Println("排序后的数组:", sortedArr)
}
