package leetcode

func palindromePartition(s string) (result [][]string) {
	if len(s) == 0 {
		return
	}

	palindromePartitionBacktracking(&result, []string{}, &s, 0)
	return
}

func palindromePartitionBacktracking(result *[][]string, curQueue []string, s *string, idx int) {
	if len(curQueue) > 0 && idx == len(*s) {
		ss := make([]string, len(curQueue))
		copy(ss, curQueue)
		*result = append(*result, ss)
		return
	}

	for i := idx; i < len(*s); i++ {
		if isPalindromePartition(s, idx, i) {
			curQueue = append(curQueue, (*s)[idx:i+1])
			palindromePartitionBacktracking(result, curQueue, s, i+1)
			curQueue = curQueue[:len(curQueue)-1]
		}
	}
}

func isPalindromePartition(s *string, lo, hi int) bool {
	if lo == hi {
		return true
	}

	for lo < hi {
		if (*s)[lo] != (*s)[hi] {
			return false
		}
		lo++
		hi--
	}

	return true
}
