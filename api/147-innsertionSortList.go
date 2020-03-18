package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	for curNode := head; curNode != nil && curNode.Next != nil; {
		if curNode.Val <= curNode.Next.Val {
			curNode = curNode.Next
			continue
		}

		prevInsert, toInsert := dummyHead, curNode.Next
		for prevInsert.Next.Val < toInsert.Val {
			prevInsert = prevInsert.Next
		}

		curNode.Next = toInsert.Next
		toInsert.Next = prevInsert.Next
		prevInsert.Next = toInsert
	}

	return dummyHead.Next
}
