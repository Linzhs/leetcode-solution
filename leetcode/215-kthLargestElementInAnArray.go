package leetcode

import "container/heap"

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 构建K个元素的最小堆
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	minHeap := &IntMinHeap{}
	for _, v := range nums {
		heap.Push(minHeap, v)
		if len(*minHeap) > k {
			heap.Pop(minHeap)
		}
	}

	return (*minHeap)[0]
}
