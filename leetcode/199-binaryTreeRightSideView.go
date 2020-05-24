package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) (result []int) {
	if root == nil {
		return
	}

	var (
		rightmostDepthNodeValMapping = make(map[int]int)
		nodeDequeue                  []*TreeNode
		treeDepths                   []int
		maxDepth                     = -1
	)
	nodeDequeue = append(nodeDequeue, root)
	treeDepths = append(treeDepths, 0)

	for len(nodeDequeue) > 0 {
		node := nodeDequeue[0]
		depth := treeDepths[0]
		nodeDequeue = nodeDequeue[1:]
		treeDepths = treeDepths[1:]

		if node == nil {
			continue
		}

		maxDepth = max(maxDepth, depth)
		rightmostDepthNodeValMapping[depth] = node.Val

		nodeDequeue = append(nodeDequeue, node.Left, node.Right)
		treeDepths = append(treeDepths, depth+1, depth+1)
	}

	for depth := 0; depth <= maxDepth; depth++ {
		result = append(result, rightmostDepthNodeValMapping[depth])
	}

	return
}
