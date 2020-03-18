package api

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		wordDictSet[v] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordDictSet[s[j:i]]; ok && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
