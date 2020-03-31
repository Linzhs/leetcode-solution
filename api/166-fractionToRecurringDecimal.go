package api

import (
	"strconv"
	"strings"
)

// 不考虑分母为0的情况
// 无限不循环小数是不能写成分数的形式的
func fractionToDecimal(numerator, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	fn := func(n int) int {
		if n > 0 {
			return 1
		}
		return 0
	}

	sb := strings.Builder{}
	if fn(numerator)^fn(denominator) == 1 {
		sb.WriteRune('-')
	}

	dividend := abs(numerator)  // 被除数
	divisor := abs(denominator) // 除数
	sb.WriteString(strconv.Itoa(int(dividend / divisor)))

	remainder := dividend % divisor
	if remainder == 0 {
		return sb.String()
	}

	sb.WriteRune('.')
	m := make(map[int64]int) // 余数 -> 位置
	for remainder > 0 {
		// 关键点 无限循环小数
		if l, ok := m[remainder]; ok {
			s := sb.String()
			return s[:l] + "(" + s[l:sb.Len()] + ")"
		}

		m[remainder] = sb.Len()
		remainder *= 10
		sb.WriteString(strconv.Itoa(int(remainder / divisor)))
		remainder %= divisor
	}

	return sb.String()
}

func abs(x int) int64 {
	if x < 0 {
		return -int64(x)
	}
	return int64(x)
}
