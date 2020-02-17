package api

import "testing"

func makeLinkedList(nums []int) (head *ListNode) {
	if len(nums) == 0 {
		return
	}

	head = &ListNode{Val: nums[0]}

	curNode := head

	for i := 1; i < len(nums); i++ {
		curNode.Next = &ListNode{Val: nums[i]}
		curNode = curNode.Next
	}

	return
}

func TestPartition(t *testing.T) {
	tests := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		{head: makeLinkedList([]int{1, 4, 3, 2, 5, 2}), x: 3, want: makeLinkedList([]int{1, 2, 2, 4, 3, 5})},
		{head: makeLinkedList([]int{1, 1}), x: 0, want: makeLinkedList([]int{1, 1})},
	}

	for _, v := range tests {
		got := partition(v.head, v.x)
		if got != v.want {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, got)
		}
	}
}
