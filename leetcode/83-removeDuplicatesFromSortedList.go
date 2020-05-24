package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for curNode := head; curNode.Next != nil; {
		if curNode.Val != curNode.Next.Val {
			curNode = curNode.Next
			continue
		}

		curNode.Next = curNode.Next.Next
	}

	return head
}
