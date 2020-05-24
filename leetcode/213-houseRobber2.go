package leetcode

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robWithRange(nums, 0, len(nums)-2), robWithRange(nums, 1, len(nums)-1))
}

func robWithRange(nums []int, lo, hi int) int {
	dp1, dp2 := 0, 0
	for i := lo; i <= hi; i++ {
		dp1, dp2 = dp2, max(dp1+nums[i], dp2)
	}

	return dp2
}
