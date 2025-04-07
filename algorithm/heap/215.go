package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 6, 1, 23, 9, 11}, 3))

}

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	fmt.Println("first final:", nums)
	// for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
	// 	nums[0], nums[i] = nums[i], nums[0]
	// 	heapSize--
	// 	maxHeapify(nums, 0, heapSize)
	// }
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	fmt.Println(a, i, heapSize)
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
