package main

import "fmt"

func heapify(arr []int, i int, length int) {
	l, r := i*2+1, i*2+2
	tag := i
	if l < length && arr[l] > arr[tag] {
		tag = l
	}
	if r < length && arr[r] > arr[tag] {
		tag = r
	}
	if tag != i {
		arr[tag], arr[i] = arr[i], arr[tag]
		heapify(arr, tag, length)
	}
}

func heapSort(arr []int) {
	// 构建最大堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	// 逐步取出堆顶元素并调整堆
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func main() {
	arr := []int{1, 8, 16, 15, 9, 12}
	fmt.Println("Unsorted array:", arr)
	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}
