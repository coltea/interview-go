package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTopK([]int{22, 6, 1, 23, 9, 25, 11}, 2))
}

func findTopK(arr []int, k int) int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		Heapify(arr, i, len(arr))
	}
	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]

	for i := len(arr) - 2; i >= len(arr)-k; i-- {
		// fmt.Println(arr)
		Heapify(arr, 0, i+1)
		arr[0], arr[i] = arr[i], arr[0]
	}
	// fmt.Println(arr)
	return arr[len(arr)-k]
}

func Heapify(arr []int, index int, length int) {

	l, r := index*2+1, index*2+2
	maxIndex := index
	if l < length && arr[l] > arr[maxIndex] {
		maxIndex = l
	}

	if r < length && arr[r] > arr[maxIndex] {
		maxIndex = r
	}

	if maxIndex != index {
		arr[maxIndex], arr[index] = arr[index], arr[maxIndex]
		Heapify(arr, maxIndex, length)
	}
}
