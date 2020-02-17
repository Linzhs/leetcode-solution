package api

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 && obstacleGrid[i][j] != 1 {
				dp[m-1][n-1] = 1
				continue
			}

			if i < m && j+1 < n && obstacleGrid[i][j+1] != 1 {
				dp[i][j] += dp[i][j+1]
			}
			if i+1 < m && j < n && obstacleGrid[i+1][j] != 1 {
				dp[i][j] += dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}
