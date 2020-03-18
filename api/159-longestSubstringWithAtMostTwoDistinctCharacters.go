package api

// 滑动窗口 使用set来记录
// eg: eceba ccaabbb
func lengthOfLongestSubstringTwoDistinct(s string) int {
	var result, winLeft int
	winSet := make(map[rune]int)

	for i, v := range s {
		winSet[v]++ // 进入窗口
		for len(winSet) > 2 {
			// 最多两个不同的字符 滑动窗口
			if cnt, ok := winSet[v]; ok && cnt != 0 {
				winSet[v]--
			}
			if winSet[v] == 0 {
				delete(winSet, v)
			}
			winLeft++
		}
		result = max(result, i-winLeft+1)
	}

	return result
}
