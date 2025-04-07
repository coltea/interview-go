package main

import (
	"fmt"
)

func findTopK(arr []int, k int) int {
	fmt.Println(arr)

	// 1. 先建一个大小为 k 的小顶堆
	for i := k/2 - 1; i >= 0; i-- {
		minHeapify(arr, i, k)
	}
	fmt.Println(arr)
	// 2. 遍历剩余元素，维护小顶堆
	for i := k; i < len(arr); i++ {
		if arr[i] > arr[0] { // 只替换比堆顶大的元素
			arr[0] = arr[i]
			minHeapify(arr, 0, k)
		}
	}

	// 3. 堆顶就是第 k 大的元素
	return arr[0]
}

func minHeapify(arr []int, index int, length int) {
	l, r := index*2+1, index*2+2
	minIndex := index

	if l < length && arr[l] < arr[minIndex] {
		minIndex = l
	}

	if r < length && arr[r] < arr[minIndex] {
		minIndex = r
	}

	if minIndex != index {
		arr[minIndex], arr[index] = arr[index], arr[minIndex]
		minHeapify(arr, minIndex, length)
	}
}

func main() {
	arr := []int{22, 6, 1, 23, 9, 25, 11}
	k := 3
	fmt.Println(findTopK(arr, k)) // 输出 5
}
