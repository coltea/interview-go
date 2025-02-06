package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
func longestConsecutive(nums []int) (ans int) {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	for i := range m {
		if m[i-1] {
			continue
		}

		i2 := i + 1
		for m[i2] {
			i2++
		}

		if i2-i > ans {
			ans = i2 - i
		}

	}
	return
}
