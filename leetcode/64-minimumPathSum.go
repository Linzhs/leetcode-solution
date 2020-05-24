package leetcode

import "math"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[m-1][n-1] = grid[m-1][n-1]
				continue
			}

			right, down := math.MaxUint32, math.MaxUint32
			if i < m && j+1 < n {
				right = dp[i][j+1]
			}
			if i+1 < m && j < n {
				down = dp[i+1][j]
			}

			switch right > down {
			case true:
				dp[i][j] = grid[i][j] + down
			case false:
				dp[i][j] = grid[i][j] + right
			}
		}
	}

	return dp[0][0]
}
