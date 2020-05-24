package leetcode

func lengthOfLongestSubstring(s string) (result int) {
	m := make(map[uint8]int)

	for i, j := 0, 0; j < len(s); j++ {
		v, ok := m[s[j]]
		if ok {
			i = max(v, i)
		}

		result = max(j-i+1, result)
		m[s[j]] = j + 1
	}

	return
}
