package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	levelOrderBottomRecursive(&result, root, 1)

	// reverse
	for lo, hi := 0, len(result)-1; lo < hi; {
		result[lo], result[hi] = result[hi], result[lo]
		lo++
		hi--
	}

	return
}

func levelOrderBottomRecursive(result *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(*result) < level {
		*result = append(*result, []int{root.Val})
	} else {
		(*result)[level-1] = append((*result)[level-1], root.Val)
	}

	levelOrderBottomRecursive(result, root.Left, level+1)
	levelOrderBottomRecursive(result, root.Right, level+1)
}
