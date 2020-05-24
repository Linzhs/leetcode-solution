package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) (result [][]int) {
	if root == nil {
		return
	}

	pathSumRecursive(&result, &[]*TreeNode{}, root, sum)

	return
}

func pathSumRecursive(result *[][]int, stack *[]*TreeNode, root *TreeNode, sum int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && sum == root.Val {
		s := make([]int, 0, len(*stack))
		for _, v := range *stack {
			s = append(s, v.Val)
		}
		s = append(s, root.Val)
		*result = append(*result, s)
		return
	}

	*stack = append(*stack, root)
	pathSumRecursive(result, stack, root.Left, sum-root.Val)
	pathSumRecursive(result, stack, root.Right, sum-root.Val)
	*stack = (*stack)[:len(*stack)-1]
}
