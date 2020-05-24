package leetcode

/**
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes 1->2->3->5.
 *
 * Note: Given n will always be valid
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	slow.Next = slow.Next.Next

	return head
}

/*
 * case:
 * 		[1] 1 ==> []
 *		[1,2] 2 ==> [2]
 */
