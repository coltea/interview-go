package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // [1,2]

	nums = []int{1}
	k = 1
	fmt.Println(topKFrequent(nums, k)) // [1]
}
