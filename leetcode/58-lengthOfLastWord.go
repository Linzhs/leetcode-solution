package leetcode

// "Hello World" ==> 5
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	l := len(s)
	for l > 0 && s[l-1] == ' ' {
		l--
	}

	cnt := 0
	for l > 0 && s[l-1] != ' ' {
		cnt++
		l--
	}

	return cnt
}
