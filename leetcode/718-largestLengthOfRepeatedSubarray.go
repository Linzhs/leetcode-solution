package leetcode

func findLength(A []int, B []int) (ans int) {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(B)+1)
	}

	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	return
}
