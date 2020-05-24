package leetcode

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	rightNode := root.Right

	newRoot := upsideDownBinaryTree(root.Left)
	newRoot.Left = rightNode
	newRoot.Right = root
	root.Left = nil
	root.Right = nil

	return newRoot
}
