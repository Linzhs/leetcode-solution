package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return (sum-root.Val == 0 && (root.Left == nil && root.Right == nil)) ||
		hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}
