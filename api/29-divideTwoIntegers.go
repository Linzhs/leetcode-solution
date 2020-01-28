package api

import "math"

// 15 - 2 * 1
// 15 - 2 * 2
// 15 - 2 * 4
// 15 - 2 * 8 < 0
//
// 7 - 2 * 1
// 7 - 2 * 2
// 7 - 2 * 4 < 0
//
// 3 - 2 * 1
// 1 < 2
func divide(dividend int, divisor int) int {
	// edge case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 符号
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	// for loop
	ans := 0
	for dividend, divisor := intAbs(dividend), intAbs(divisor); dividend >= divisor; {
		curSht, curDivisor := 1, divisor

		for dividend >= (curDivisor << 1) {
			curDivisor <<= 1
			curSht <<= 1
		}

		dividend -= curDivisor
		ans += curSht
	}

	return sign * ans
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
