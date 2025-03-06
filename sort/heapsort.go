package main

import "fmt"

// heapify 维护最大堆性质
func heapify(arr []int, n, i int) {
	fmt.Println(arr, n, i)
	largest := i
	left, right := 2*i+1, 2*i+2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// HeapSort 执行堆排序
func HeapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	fmt.Println(111, arr)
	// 依次取出堆顶元素，并调整堆
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{1, 8, 3, 15, 9, 12}
	fmt.Println("Unsorted array:", arr)
	HeapSort(arr)
	fmt.Println("Sorted array:", arr)

	// arr = []int{4, 78, 1, 5, 1236, 0}
	// fmt.Println("Unsorted array:", arr)
	// HeapSort(arr)
	// fmt.Println("Sorted array:", arr)

}
