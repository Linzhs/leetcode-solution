package leetcode

// You may assume both s and t have the same length.
// paper -> title or title -> paper
func isIsomorphic(s string, t string) bool {
	if len(s) <= 1 && len(t) <= 1 {
		return true
	}

	return checkIsomorphic(s, t) && checkIsomorphic(t, s)
}

func checkIsomorphic(s, t string) bool {
	m := make(map[byte]byte)
	for i, v := range []byte(s) {
		if a, ok := m[v]; ok {
			if a != t[i] {
				return false
			}
			continue
		}

		m[v] = t[i]
	}

	return true
}
