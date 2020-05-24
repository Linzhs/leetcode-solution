package leetcode

import "math"

func numSquares(n int) int {
	var squareNums []int
	for i := 1; i < int(math.Sqrt(float64(n))+1); i++ {
		squareNums = append(squareNums, i*i)
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i <= n; i++ {
		for _, v := range squareNums {
			if i < v {
				break
			}
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}

	return dp[n]
}
