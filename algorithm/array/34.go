package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	fmt.Println(searchRange([]int{}, 7))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
	fmt.Println(searchRange([]int{1}, 1))

}

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

//
//
// func searchRange(nums []int, target int) []int {
// 	index := BinarySearch(nums, target)
// 	if index == -1 {
// 		return []int{-1, -1}
// 	}
// 	s, e := index, index
// 	for s-1 >= 0 && nums[s-1] == target {
// 		s--
// 	}
// 	for e+1 <= len(nums)-1 && nums[e+1] == target {
// 		e++
// 	}
// 	return []int{s, e}
//
// }

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // 防止溢出
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
