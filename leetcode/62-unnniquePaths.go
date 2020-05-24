package leetcode

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化状态转移方程
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[m-1][n-1] = 1
				continue
			}

			if i < m && j+1 < n {
				dp[i][j] += dp[i][j+1]
			}
			if i+1 < m && j < n {
				dp[i][j] += dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}
