package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// note: 1 ≤ m ≤ n ≤ length of list.
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	var (
		reverseBeforeNode *ListNode
		reversedEndNode   *ListNode
		prev              *ListNode
		curr              = head
	)

	for i := 1; i <= n; i++ {
		if i < m {
			reverseBeforeNode, curr = curr, curr.Next
			continue
		}

		// m <= i <= n
		if i == m {
			reversedEndNode = curr
		}

		// mark next node
		next := curr.Next

		// reverse
		curr.Next = prev

		// forward
		prev, curr = curr, next
	}

	reversedEndNode.Next = curr
	if reverseBeforeNode != nil {
		reverseBeforeNode.Next = prev
	}

	if m == 1 {
		return prev
	}

	return head
}
