package api

import "math"

func rangeBitwiseAnd(m int, n int) (result int) {
	result = m
	for i := m + 1; i <= n; i++ {
		result &= i
		if result == 0 || i == math.MaxInt32 {
			break
		}
	}
	return
}
