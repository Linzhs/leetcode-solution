package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	newRoot := &TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}

	return newRoot
}
