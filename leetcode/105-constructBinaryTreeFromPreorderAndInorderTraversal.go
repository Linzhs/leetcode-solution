package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	rootPos := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootPos = i
			break
		}
	}

	root.Left = buildTree(preorder[1:rootPos+1], inorder[:rootPos])
	root.Right = buildTree(preorder[rootPos+1:], inorder[rootPos+1:])

	return root
}
