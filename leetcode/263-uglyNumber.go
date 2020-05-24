package leetcode

func isUgly(num int) bool {
	if num == 1 {
		return true
	}

	for num > 1 && (num%2 == 0 || num%3 == 0 || num%5 == 0) {
		if num%2 == 0 {
			num /= 2
		}
		if num%3 == 0 {
			num /= 3
		}
		if num%5 == 0 {
			num /= 5
		}
	}

	return num == 1
}
