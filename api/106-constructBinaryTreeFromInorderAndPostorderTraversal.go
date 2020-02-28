package api

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	rootPos := 0
	for i, v := range inorder {
		if v == root.Val {
			rootPos = i
			break
		}
	}

	root.Left = buildTree2(inorder[:rootPos], postorder[:rootPos])
	root.Right = buildTree2(inorder[rootPos+1:], postorder[rootPos:len(postorder)-1])

	return root
}
