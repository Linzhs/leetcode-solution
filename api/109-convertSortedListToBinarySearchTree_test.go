package api

import "testing"

func TestSortedListToBST(t *testing.T) {
	sortedListToBST(makeLinkedList([]int{-10, -3, 0, 5, 9}))
}
