package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return sortedArrayToBstRecursive(nums, 0, len(nums)-1)
}

func sortedArrayToBstRecursive(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	mid := (lo + hi) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBstRecursive(nums, lo, mid-1),
		Right: sortedArrayToBstRecursive(nums, mid+1, hi),
	}

	return root
}
