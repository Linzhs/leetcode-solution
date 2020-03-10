package api

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) (sum int) {
	return sumNumbersRecursive(root, 0)
}

func sumNumbersRecursive(root *TreeNode, curVal int) int {
	if root == nil {
		return 0
	}

	curVal = curVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return curVal
	}

	return sumNumbersRecursive(root.Left, curVal) + sumNumbersRecursive(root.Right, curVal)
}
