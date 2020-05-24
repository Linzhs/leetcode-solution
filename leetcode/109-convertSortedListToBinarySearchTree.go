package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// [-10,-3,0,5,9]
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	tailNodePosition := -1
	for cur := head; cur != nil; cur = cur.Next {
		tailNodePosition++
	}

	return sortedListToBstRecursive(head, 0, tailNodePosition)
}

// 可用中序遍历的思路提升速度
func sortedListToBstRecursive(head *ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	cur := head
	for i := 0; i < mid; i++ {
		cur = cur.Next
	}

	return &TreeNode{
		Val:   cur.Val,
		Left:  sortedListToBstRecursive(head, start, mid-1),
		Right: sortedListToBstRecursive(head, mid+1, end),
	}
}
