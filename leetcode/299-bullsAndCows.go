package leetcode

import "fmt"

func getHint(secret, guess string) string {
	var bulls, cows int
	sNumbers := make([]int, 10)
	gNumbers := make([]int, 10)
	for i, s := range secret {
		g := guess[i]
		if s == int32(g) {
			bulls++
			continue
		}

		sNumbers[s-'0']++
		gNumbers[g-'0']++
	}

	for i := 0; i < 10; i++ {
		cows += min(sNumbers[i], gNumbers[i])
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
