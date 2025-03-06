package main

import "fmt"

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort2(arr[:left])
	quickSort2(arr[left+1:])

	return arr
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Unsorted array:", arr)
	sortedArr := quickSort2(arr)
	fmt.Println("Sorted array:", sortedArr)
}
