package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{1, 3, 4, 5, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		}
		if sum < target {
			i++
		}

	}

	return []int{0, 0}
}
