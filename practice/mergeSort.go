package main

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {

	res := make([]int, 0, len(left)+len(right))
	i1, i2 := 0, 0
	for i1 < len(left) && i2 < len(right) {
		if left[i1] <= right[i2] {
			res = append(res, left[i1])
			i1++
		} else {
			res = append(res, right[i2])
			i2++
		}
	}

	res = append(res, left[i1:]...)
	res = append(res, right[i2:]...)

	return res

}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted:", arr)

	sortedArr := MergeSort(arr)
	fmt.Println("Sorted:", sortedArr)
}
