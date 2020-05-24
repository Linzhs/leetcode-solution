package leetcode

func maximalSquare(matrix [][]byte) int {
	const charOne = '1'

	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}

	var (
		dp    = make([][]int, m+1)
		sqLen = 0
	)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] != charOne {
				continue
			}

			dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
			if sqLen < dp[i][j] {
				sqLen = dp[i][j]
			}
		}
	}

	return sqLen * sqLen
}
