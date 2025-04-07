package main

import (
	"fmt"
)

func main() {
	// 测试归并排序
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("原始数组:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("排序后的数组:", sortedArr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	s1 := arr[:len(arr)/2]
	s2 := arr[len(arr)/2:]
	return merge(mergeSort(s1), mergeSort(s2))
}

func merge(s1, s2 []int) []int {
	res := make([]int, len(s1)+len(s2))
	l, r := 0, 0
	for l < len(s1) && r < len(s2) {
		if s1[l] < s2[r] {
			res = append(res, s1[l])
			l++
		} else {
			res = append(res, s2[r])
			r++
		}
	}
	res = append(res, s1[l:]...)
	res = append(res, s2[r:]...)
	return res
}
