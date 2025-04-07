package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 8, 16, 15, 9, 12}
	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr)/2 - 1
	fmt.Println("arr:", arr)
	arr2 := append(arr[:mid], arr[mid+1:]...)
	fmt.Println("arr2:", arr2)
	l, r := make([]int, 0), make([]int, 0)
	for i, n := range arr2 {
		if n <= arr[mid] {
			l = append(l, arr2[i])
		} else {
			r = append(r, arr2[i])
		}
	}

	return append(append(quickSort(l), arr[mid]), quickSort(r)...)
}
