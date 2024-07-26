package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		sub := target - num
		if _, ok := m[sub]; !ok {
			m[sub] = i
		}
	}
	// fmt.Println(m)
	for i, num := range nums {
		if j, ok := m[num]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
