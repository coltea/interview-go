package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topK(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	return *h
}

func main() {
	nums := []int{3, 1, 71, 12, 2, 11, 171, 89, 9}
	k := 3
	fmt.Printf("Top %d elements: %v\n", k, topK(nums, k))
}
