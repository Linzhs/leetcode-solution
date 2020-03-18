package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 快慢指针找到链表中间节点
	prev, slow, fast := (*ListNode)(nil), head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // 断开前后链表

	// 反转链表后半段
	prev, cur, next := (*ListNode)(nil), slow, slow.Next
	for next != nil {
		cur.Next = prev

		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev

	// 重排
	// 1->2
	// 5->4->3
	//
	// 1->2->3
	// 6->5->4
	p1, p2 := head, cur
	for p2 != nil {
		next1, next2 := p1.Next, p2.Next

		p1.Next = p2

		if next1 == nil {
			break
		}

		p2.Next = next1

		p1, p2 = next1, next2
	}
}
