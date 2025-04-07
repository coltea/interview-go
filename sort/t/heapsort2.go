package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 8, 16, 15, 9, 12}
	fmt.Println("Unsorted array:", arr)
	heapSort2(arr)
	fmt.Println("Sorted array:", arr)
}

func heapSort2(arr []int) {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		heapify2(arr, i, length)
	}
	fmt.Println(arr)
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify2(arr, 0, i)
	}

}

func heapify2(arr []int, i int, length int) {
	target := i
	l, r := 2*i+1, 2*i+2

	if l < length && arr[l] > arr[target] {
		target = l
	}

	if r < length && arr[r] > arr[target] {
		target = r
	}
	if target != i {
		arr[target], arr[i] = arr[i], arr[target]
		heapify2(arr, target, length)
	}
}
