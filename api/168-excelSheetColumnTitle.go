package api

func convertToTitle(n int) (s string) {
	for n > 0 {
		n--
		s = string(n%26+'A') + s
		n /= 26
	}

	return
}
