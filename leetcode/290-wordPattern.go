package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	splitString := strings.Split(str, " ")

	if len(splitString) != len(pattern) {
		return false
	}

	var (
		valueSet = make(map[string]struct{})
		hm       = make(map[rune]string)
	)

	for i, v := range pattern {
		word := splitString[i]
		s, ok1 := hm[v]
		if !ok1 {
			_, ok2 := valueSet[word]
			if ok2 {
				return false
			}
			hm[v] = word
			valueSet[word] = struct{}{}
			continue
		}

		if s != word {
			return false
		}
	}

	return true
}
