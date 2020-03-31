package api

func rob(nums []int) int {
	dp1, dp2 := 0, 0
	for _, v := range nums {
		dp1, dp2 = dp2, max(dp1+v, dp2)
	}
	return dp2
}
