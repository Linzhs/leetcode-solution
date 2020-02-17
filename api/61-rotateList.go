package api

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var nodeNums int
	tailNode := head
	for curNode := head; curNode != nil; curNode = curNode.Next {
		nodeNums++
		tailNode = curNode
	}

	rotateSub := k % nodeNums
	if rotateSub == 0 {
		return head
	}

	for i, curNode := 1, head; curNode.Next != nil; curNode = curNode.Next {
		if i == nodeNums-rotateSub {
			nextNode := curNode.Next
			tailNode.Next = head
			curNode.Next = nil
			return nextNode
		}

		i++
	}

	return head
}
