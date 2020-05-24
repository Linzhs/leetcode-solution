package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	var stack1, stack2 []*TreeNode
	stack1 = append(stack1, root)

	fromLeft := false

	for len(stack1) != 0 || len(stack2) != 0 {
		if len(stack1) == 0 || len(stack2) == 0 {
			// 只要有一个栈为空就反转入栈方向
			fromLeft = !fromLeft
		}

		if len(stack1) != 0 {
			var levels []int
			for i := len(stack1) - 1; i >= 0; i-- {
				levels = append(levels, stack1[i].Val)
				if fromLeft {
					if stack1[i].Left != nil {
						stack2 = append(stack2, stack1[i].Left)
					}
					if stack1[i].Right != nil {
						stack2 = append(stack2, stack1[i].Right)
					}
				} else {
					if stack1[i].Right != nil {
						stack2 = append(stack2, stack1[i].Right)
					}
					if stack1[i].Left != nil {
						stack2 = append(stack2, stack1[i].Left)
					}
				}
			}
			result = append(result, levels)
			stack1 = []*TreeNode{}
			continue
		}

		if len(stack2) != 0 {
			var levels []int
			for i := len(stack2) - 1; i >= 0; i-- {
				levels = append(levels, stack2[i].Val)
				if fromLeft {
					if stack2[i].Left != nil {
						stack1 = append(stack1, stack2[i].Left)
					}
					if stack2[i].Right != nil {
						stack1 = append(stack1, stack2[i].Right)
					}
				} else {
					if stack2[i].Right != nil {
						stack1 = append(stack1, stack2[i].Right)
					}
					if stack2[i].Left != nil {
						stack1 = append(stack1, stack2[i].Left)
					}
				}
			}
			result = append(result, levels)
			stack2 = []*TreeNode{}
		}
	}

	return
}
