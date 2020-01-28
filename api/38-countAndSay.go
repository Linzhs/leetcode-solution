package api

import (
	"strconv"
)

//	1.     1
//	2.     11
//	3.     21
//	4.     1211
//	5.     111221
func countAndSay(n int) string {
	curSay := "1"
	if n == 1 {
		return curSay
	}

	for i := 2; i <= n; i++ {
		var (
			count  int
			say    rune
			curRes string
		)

		for _, v := range curSay {
			switch {
			case say == 0:
				count++
				say = v
			case say == v:
				count++
			case say != v:
				curRes += strconv.Itoa(count) + string(say)
				count = 1
				say = v
			}
		}

		if say != 0 {
			curRes += strconv.Itoa(count) + string(say)
		}

		curSay = curRes
	}

	return curSay
}
