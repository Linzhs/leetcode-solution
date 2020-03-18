package api

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	l := makeLinkedList([]int{1, 2, 3, 4, 5})
	reorderList(l)

	for cur := l; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
