package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找出链表中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid, midPrev := slow.Next, slow

	// 反转后半端节点
	halfHead := reverseLinkedList(mid)
	result := true
	for p1, p2 := head, halfHead; result && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			result = false
		}
	}

	// 还原后半段链表
	midPrev.Next = reverseLinkedList(mid)

	return result
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr, next := (*ListNode)(nil), head, head.Next
	for next != nil {
		curr.Next = prev

		prev = curr
		curr = next
		next = next.Next
	}
	curr.Next = prev

	return curr
}
