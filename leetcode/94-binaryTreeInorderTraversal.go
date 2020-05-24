package leetcode

import "leetcode-solution/pkg"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *pkg.TreeNode) []int {

	var result []int

	var stack []*pkg.TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = top.Right

			result = append(result, top.Val)
		}
	}

	return result
}
