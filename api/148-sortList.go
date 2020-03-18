package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 自底向上
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 统计链表长度
	listLen := 0
	for cur := head; cur != nil; cur = cur.Next {
		listLen++
	}

	dummyHead := &ListNode{Next: head}
	for step := 1; step < listLen; step <<= 1 {
		curNode := dummyHead.Next
		prevTail := dummyHead
		for curNode != nil {
			left := curNode
			right := splitLinkedList(left, step)
			curNode = splitLinkedList(right, step)
			prevTail = merge2SortedLinkedList(left, right, prevTail)
		}
	}

	return dummyHead.Next
}

// splitLinkedList 分裂指定长度链表 返回分裂完成后另一端的head节点
func splitLinkedList(head *ListNode, n int) *ListNode {
	curNode := head
	for i := 1; curNode != nil && i < n; i++ {
		curNode = curNode.Next
	}

	if curNode == nil {
		return nil
	}

	nextHead := curNode.Next
	curNode.Next = nil
	return nextHead
}

// merge2SortedLinkedList 合并两条链表(完成排序)
// 返回尾指针 以便下一次迭代链接另一条链表
func merge2SortedLinkedList(l1, l2, head *ListNode) *ListNode {
	curNode := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curNode.Next = l2
			curNode = l2
			l2 = l2.Next
			continue
		}

		curNode.Next = l1
		curNode = l1
		l1 = l1.Next
	}

	// 还有剩余
	if l1 != nil {
		curNode.Next = l1
	} else {
		curNode.Next = l2
	}

	// 指向尾端
	for curNode.Next != nil {
		curNode = curNode.Next
	}

	return curNode
}
