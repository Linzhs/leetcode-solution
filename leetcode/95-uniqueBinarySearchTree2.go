package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	// dp从1到n
	// 状态转移方程为
	//		1. n为根节点 dp[n-1]的所有元素为其左子树
	//		2. n为非根节点 向dp[n-1]的所有树的右子树插入n
	// 由于生成当前dp只需要n-1的dp 所以只需要两个slice 这里设置为dp[0~1]
	dp := make([][]*TreeNode, 2)
	dp[0] = []*TreeNode{{
		Val:   1,
		Left:  nil,
		Right: nil,
	}}

	for i := 2; i <= n; i++ {
		for _, root := range dp[0] {
			// n为根节点的情况 只有一种
			dp[1] = append(dp[1], &TreeNode{
				Val:   i,
				Left:  treeClone(root),
				Right: nil,
			})

			// n为非根节点的情况 需要向该树的右子树迭代的插入n <只向右不向左 元素val唯一>
			for curNode := root; curNode != nil; curNode = curNode.Right {
				// 先把树克隆出一份
				newRoot := treeClone(root)

				// 找到当前遍历到的节点
				for node := newRoot; node != nil; node = node.Right {
					if node.Val == curNode.Val {
						node.Right = &TreeNode{
							Val:  i,
							Left: treeClone(node.Right),
						}
						break
					}
				}

				dp[1] = append(dp[1], newRoot)
			}
		}

		dp[0] = dp[1]
		dp[1] = []*TreeNode{}
	}

	return dp[0]
}

func treeClone(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  treeClone(root.Left),
		Right: treeClone(root.Right),
	}
}
