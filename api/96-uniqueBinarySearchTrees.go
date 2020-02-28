package api

// i ==> left: 1~i-1 ; right: i+1~n
// i是递增的 长度相同的子树排出不同情况个数相同  所以只和长度有关
func numTrees(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j] // 左子树可能的情况 * 右子树可能的情况
		}
	}

	return dp[n]
}
