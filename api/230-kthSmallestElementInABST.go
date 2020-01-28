package api

import "leetcode-solution/pkg"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *pkg.TreeNode, k int) int {
	// 中序遍历有序
	cur := root
	var stack []*pkg.TreeNode
	var kthSlice []int
	for cur != nil || len(stack) != 0 {
		switch cur != nil {
		case true:
			stack = append(stack, cur)
			cur = cur.Left
		default:
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = top.Right

			kthSlice = append(kthSlice, top.Val)
			if len(kthSlice) == k {
				return top.Val
			}
		}
	}
	return 0
}
