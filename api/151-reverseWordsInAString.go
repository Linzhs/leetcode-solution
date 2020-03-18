package api

// 思路:
// 1. 反转整个字符串
// 2. 反转每个单词字符串
// 3. 清除空格
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	stringBytes := []byte(s)

	reverseStringWithInterval(stringBytes, 0, len(s)-1)

	for i, j := 0, 0; i < len(stringBytes); {
		// 找到单词左边界
		for i < j || i < len(stringBytes) && stringBytes[i] == ' ' {
			i++
		}
		// 找到单词右边界
		for j < i || j < len(stringBytes) && stringBytes[j] != ' ' {
			j++
		}
		reverseStringWithInterval(stringBytes, i, j-1)
	}

	return string(cleanSpace(stringBytes))
}

func reverseStringWithInterval(s []byte, lo, hi int) {
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}

func cleanSpace(s []byte) []byte {
	i, j := 0, 0
	for j < len(s) {
		for j < len(s) && s[j] == ' ' {
			j++
		}
		for j < len(s) && s[j] != ' ' {
			s[i] = s[j]
			i++
			j++
		}
		for j < len(s) && s[j] == ' ' {
			j++
		}
		if j < len(s) {
			s[i] = ' '
			i++
		}
	}

	return s[:i]
}
