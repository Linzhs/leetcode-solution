package api

func titleToNumber(s string) (result int) {
	for _, c := range s {
		result = result*26 + int(c-'A'+1)
	}
	return
}
