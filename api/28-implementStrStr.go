package api

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	} else if len(needle) > len(haystack) {
		return -1
	}

	sBytes := []byte(haystack)
	nBytes := []byte(needle)

	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(nBytes) {
				return i
			}
			if i+j == len(sBytes) {
				return -1
			}
			if sBytes[i+j] != nBytes[j] {
				break
			}
		}
	}
}
