package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pHead1, pHead2, curP1, curP2 *ListNode
	for curNode := head; curNode != nil; curNode = curNode.Next {
		switch curNode.Val < x {
		case true:
			if pHead1 == nil {
				pHead1 = curNode
				curP1 = curNode
			} else {
				curP1.Next = curNode
				curP1 = curNode
			}
		case false:
			if pHead2 == nil {
				pHead2 = curNode
				curP2 = curNode
			} else {
				curP2.Next = curNode
				curP2 = curNode
			}
		}
	}

	if pHead1 == nil {
		curP2.Next = nil
		return pHead2
	}

	if pHead2 == nil {
		curP1.Next = nil
		return pHead1
	}

	curP2.Next = nil
	curP1.Next = pHead2

	return pHead1
}
