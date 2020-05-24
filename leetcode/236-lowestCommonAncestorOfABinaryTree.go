package leetcode

import "leetcode-solution/pkg"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *pkg.TreeNode) *pkg.TreeNode {

	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
			continue
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
			continue
		}
		return cur
	}
	return nil
}
