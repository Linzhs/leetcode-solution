package leetcode

func trailingZeroes(n int) (result int) {
	for n > 0 {
		n /= 5
		result += n
	}
	return
}
