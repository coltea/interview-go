package main

import (
	"fmt"
)

func main() {
	a := []int{1, 433, 7, 43, 23, 12}
	fmt.Println(quickSort(a))
}

func quickSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	left, right := make([]int, 0), make([]int, 0)
	m := a[l/2]
	arr := append(a[:l/2], a[l/2+1:]...)
	for _, i := range arr {
		if i <= m {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}
	return append(append(quickSort(left), m), quickSort(right)...)
}
