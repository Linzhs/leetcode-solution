package api

import "unicode"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	for lo, hi := 0, len(s)-1; lo < hi; {
		for lo < hi && isNotAlphanumeric(s[lo]) {
			lo++
		}

		for lo < hi && isNotAlphanumeric(s[hi]) {
			hi--
		}

		if unicode.ToLower(rune(s[lo])) != unicode.ToLower(rune(s[hi])) {
			return false
		}
		lo++
		hi--
	}

	return true
}

func isNotAlphanumeric(r uint8) bool {
	if r < '0' || (r > '9' && r < 'A') || (r > 'Z' && r < 'a') || r > 'z' {
		return true
	}

	return false
}
