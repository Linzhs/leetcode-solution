package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 第一步 通过快慢指针找到相遇点
	var p *ListNode
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p = slow
			break
		}
	}

	if p == nil {
		return nil
	}

	// 第二步 找到环的入口点
	p1, p2 := head, p
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
