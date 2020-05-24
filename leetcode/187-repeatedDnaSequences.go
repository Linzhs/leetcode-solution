package leetcode

// 只有A、C、G、T四个字符
func findRepeatedDnaSequences(s string) (result []string) {
	// 二进制编码表
	// 00 -> A
	// 01 -> C
	// 02 -> G
	// 03 -> T
	chars := make([]int, 26)
	chars['C'-'A'] = 1
	chars['G'-'A'] = 2
	chars['T'-'A'] = 3

	wordSet := make(map[int]int)
	for i := 0; i < len(s)-9; i++ {
		v := 0
		// 对每10个连续字符进行二进制编码
		for j := i; j < i+10; j++ {
			v <<= 2
			v |= chars[s[j]-'A']
		}

		// 编码后的字符串是否出现过两次
		cnt, ok := wordSet[v]
		if !ok {
			// 没出现过
			wordSet[v] = 1
			continue
		}
		// 出现过 但是只出现一次
		if cnt == 1 {
			wordSet[v]++
			result = append(result, s[i:i+10])
			continue
		}
		// 出现过 已出现两次
	}

	return
}
