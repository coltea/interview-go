package main

import (
	"fmt"
)

func main() {
	a := []int{1, 433, 7, 43, 23, 12}
	fmt.Println(quickSort(a))
}

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	l, r := make([]int, 0), make([]int, 0)

	m := a[len(a)/2]
	a = append(a[:len(a)/2], a[len(a)/2+1:]...)
	for v := range a {
		if a[v] <= m {
			l = append(l, a[v])
		} else {
			r = append(r, a[v])
		}
	}

	return append(append(quickSort(l), m), quickSort(r)...)
}
