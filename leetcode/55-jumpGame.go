package leetcode

// 动态规划方法
//  dp[i] 定义为为此position是GOOD index， 可以跳到最后一个index
//			默认为unknown 无法跳到最后一个index为bad
//  dp[i] 的状态转义方程为 nums[i]+{0...position} 是否有 GOOD
const (
	unknown = iota
	good
	bad
)

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	// initial
	dp := make([]int, len(nums))
	dp[len(dp)-1] = good

	for i := len(nums) - 2; i >= 0; i-- {
		furthestPosition := min(len(nums)-1, i+nums[i])
		for j := i + 1; j <= furthestPosition; j++ {
			if dp[j] == good {
				dp[i] = good
				break
			}
		}
	}

	return dp[0] == good
}

func canJumpBackTracking(nums []int) bool {
	return canJumpFromPosition(nums, 0)
}

func canJumpFromPosition(nums []int, position int) bool {
	if position == len(nums)-1 {
		return true
	}

	furthestPosition := min(len(nums)-1, position+nums[position])
	for i := position + 1; i <= furthestPosition; i++ {
		if canJumpFromPosition(nums, i) {
			return true
		}
	}

	return false
}
