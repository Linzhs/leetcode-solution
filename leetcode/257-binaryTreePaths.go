package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) (result []string) {
	if root == nil {
		return
	}

	binaryTreePathsRecursive(&result, root, "")

	return
}

func binaryTreePathsRecursive(result *[]string, root *TreeNode, curPath string) {
	if root == nil {
		return
	}

	curPath += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*result = append(*result, curPath)
		return
	}

	curPath += "->"
	binaryTreePathsRecursive(result, root.Left, curPath)
	binaryTreePathsRecursive(result, root.Right, curPath)
}
