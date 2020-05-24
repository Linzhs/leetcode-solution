package leetcode

import "math"

func isOneEditDistance(s, t string) bool {
	sub := math.Abs(float64(len(s) - len(t)))
	switch sub {
	case 0:
		diffCnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				diffCnt++
			}
		}
		return diffCnt == 1
	case 1:
		for i := 0; i < len(t); i++ {
			if s[i] != t[i] {
				return s[i+1:] == t[i+1:]
			}
		}
		return true
	}

	return false
}
