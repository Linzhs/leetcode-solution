package api

func maxSubArray(nums []int) int {
	return maxSubArrayByDynamicProgram(nums)
}

// 最大连续子序列和 使用动态规划方式实现
// dp[i] 表示以nums[i]为结尾element的最大连续子序列和
// 由dp状态定义可得dp[i]存在两种场景
// 		==> 1. dp[i]为当前element 即nums[i]单个元素为0~i之间最大的连续子序列和
//		==> 2. dp[i]为nums[p](p<i)到nums[i]之间连续序列之和 且在0~i之间最大的连续子序列和
// 两种场景的状态为
// 		==> 1. dp[i] = nums[i]
//		==> 2. dp[i] = dp[i-1] + nums[i]
// 可得出状态转移方程为
//		==> dp[i] = max{nums[i], dp[i-1] + nums[i]}
func maxSubArrayByDynamicProgram(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(dp[i], maxSum)
	}

	return maxSum
}

// 最大连续子序列和 使用分治法
func maxSubArrayByDivideConquer(nums []int) int {
	return maxSubArrayRecursive(nums, 0, len(nums)-1)
}

func maxSubArrayRecursive(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := (lo + hi) / 2
	left := maxSubArrayRecursive(nums, lo, mid)
	right := maxSubArrayRecursive(nums, mid+1, hi)
	cross := maxSubArrayCrossSum(nums, lo, hi, mid)

	return max(max(left, right), cross)
}

func maxSubArrayCrossSum(nums []int, lo, hi, mid int) int {
	if lo == hi {
		return nums[lo]
	}

	leftMaxSum := nums[mid]
	for i := mid; i >= lo; i-- {
		leftMaxSum = max(leftMaxSum, leftMaxSum+nums[i])
	}

	rightMaxSum := nums[mid]
	for i := mid; i <= hi; i++ {
		rightMaxSum = max(rightMaxSum, rightMaxSum+nums[i])
	}

	return leftMaxSum + rightMaxSum
}
