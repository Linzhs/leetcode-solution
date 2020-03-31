package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	for prev, curNode := dummyHead, head; curNode != nil; curNode = curNode.Next {
		if curNode.Val == val {
			prev.Next = curNode.Next
			continue
		}
		prev = curNode
	}

	return dummyHead.Next
}
